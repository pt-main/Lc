package profiler

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pt-main/lc/v2"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/tooling/bytecode"
	"github.com/pt-main/lc/v2/tooling/debugging/extensiblePlugin"
	"github.com/pt-main/lc/v2/tooling/plugin"
)

const Name = "profiler"

// scopeStartTime is the key the pre hook stores the command start under in the
// shared scope, where the post hook reads it back.
const scopeStartTime = "profiler_start_time"

type Metric struct {
	Count     int64
	TotalTime time.Duration
	MinTime   time.Duration
	MaxTime   time.Duration
}

func (m *Metric) Avg() time.Duration {
	if m.Count == 0 {
		return 0
	}
	return m.TotalTime / time.Duration(m.Count)
}

// ProfilerPlugin counts and times every command through the hooks of
// ExtensibleCLPlugin, which therefore has to be installed first.
//
// Call("report") returns the collected metrics, Call("reset") clears them and
// Call("enable") / Call("disable") switch collecting on and off.
type ProfilerPlugin struct {
	plugin.Plugin
	mu               sync.Mutex
	byteMetrics      map[int]*Metric
	stringMetrics    map[string]*Metric
	totalByteCalls   int64
	totalStringCalls int64
	totalByteTime    time.Duration
	totalStringTime  time.Duration
	startTime        time.Time
	enabled          bool
	scope            core.ScopeType
}

func New() *ProfilerPlugin {
	return &ProfilerPlugin{
		byteMetrics:   make(map[int]*Metric),
		stringMetrics: make(map[string]*Metric),
		startTime:     time.Now(),
		enabled:       true,
	}
}

func (p *ProfilerPlugin) Name() string {
	return Name
}

func (p *ProfilerPlugin) Init(scope core.ScopeType, pm *plugin.PluginManager) error {
	if !(&plugin.Tools{Pm: pm}).IsPluginInstalled(extensiblePlugin.Name) {
		return fmt.Errorf("Extensible plugin is not installed, can't init profiler")
	}
	p.scope = scope
	eu, ok := scope[public.PluginsScopeEuPtr].(*lc.EngineUniversal)
	if !ok {
		return fmt.Errorf("profiler: can't get EngineUniversal")
	}
	uep, err := eu.GetUEP()
	if err != nil {
		return err
	}
	ep := extensiblePlugin.New(eu)
	err = ep.Init(scope, pm)
	if err != nil {
		return err
	}
	uep.Event.NewEvent(extensiblePlugin.CLEInPreEvent, p.preEvent)
	uep.Event.NewEvent(extensiblePlugin.CLEInPostEvent, p.postEvent)
	return nil
}

func (p *ProfilerPlugin) isEnabled() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.enabled
}

func (p *ProfilerPlugin) preEvent(ev *core.Events, _ *core.EventInput) core.ErrorInterface {
	if !p.isEnabled() {
		return nil
	}
	ev.Scope()[scopeStartTime] = time.Now()
	return nil
}

func (p *ProfilerPlugin) postEvent(ev *core.Events, _ *core.EventInput) core.ErrorInterface {
	if !p.isEnabled() {
		return nil
	}
	startVal, ok := ev.Scope()[scopeStartTime]
	if !ok {
		return nil
	}
	start, ok := startVal.(time.Time)
	if !ok {
		return nil
	}
	elapsed := time.Since(start)

	p.mu.Lock()
	defer p.mu.Unlock()

	if data, err := core.ScopeGet[extensiblePlugin.BCLEData](ev.Scope(), extensiblePlugin.CLEScopeData); err == nil {
		p.totalByteCalls++
		p.totalByteTime += elapsed
		idx := data.Idx
		if idx != nil && *idx >= 0 && *idx < len(data.Parsed) {
			attr := data.Parsed[*idx]
			opcode := 0
			if attr.RawNode != nil {
				opcode = (&bytecode.Utils{}).BytesToInt(attr.RawNode.Switch, public.LittleEndian)
			}
			p.updateMetricByte(opcode, elapsed)
		}
		return nil
	}
	if data, err := core.ScopeGet[extensiblePlugin.SCLEData](ev.Scope(), extensiblePlugin.CLEScopeData); err == nil {
		p.totalStringCalls++
		p.totalStringTime += elapsed
		idx := data.Idx
		if idx != nil && *idx >= 0 && *idx < len(data.Parsed) {
			p.updateMetricString(data.Parsed[*idx].Switch, elapsed)
		}
	}
	return nil
}

func (p *ProfilerPlugin) updateMetricByte(opcode int, elapsed time.Duration) {
	m, ok := p.byteMetrics[opcode]
	if !ok {
		m = &Metric{MinTime: elapsed, MaxTime: elapsed}
		p.byteMetrics[opcode] = m
	}
	m.Count++
	m.TotalTime += elapsed
	if elapsed < m.MinTime {
		m.MinTime = elapsed
	}
	if elapsed > m.MaxTime {
		m.MaxTime = elapsed
	}
}

func (p *ProfilerPlugin) updateMetricString(cmd string, elapsed time.Duration) {
	m, ok := p.stringMetrics[cmd]
	if !ok {
		m = &Metric{MinTime: elapsed, MaxTime: elapsed}
		p.stringMetrics[cmd] = m
	}
	m.Count++
	m.TotalTime += elapsed
	if elapsed < m.MinTime {
		m.MinTime = elapsed
	}
	if elapsed > m.MaxTime {
		m.MaxTime = elapsed
	}
}

func (p *ProfilerPlugin) Report() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	elapsed := time.Since(p.startTime)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Profiler report (%.2f sec total):\n", elapsed.Seconds()))

	if len(p.byteMetrics) > 0 {
		sb.WriteString("\n  Byte opcodes:\n")
		sb.WriteString(fmt.Sprintf("  Byte calls: %d (total time: %v)\n", p.totalByteCalls, p.totalByteTime))
		var keys []int
		for k := range p.byteMetrics {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, op := range keys {
			m := p.byteMetrics[op]
			sb.WriteString(fmt.Sprintf("    %d: count=%d, total=%v, avg=%v, min=%v, max=%v\n",
				op, m.Count, m.TotalTime, m.Avg(), m.MinTime, m.MaxTime))
		}
	}
	if len(p.stringMetrics) > 0 {
		sb.WriteString("\n  String commands:\n")
		sb.WriteString(fmt.Sprintf("  String calls: %d (total time: %v)\n", p.totalStringCalls, p.totalStringTime))
		var keys []string
		for k := range p.stringMetrics {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, cmd := range keys {
			m := p.stringMetrics[cmd]
			sb.WriteString(fmt.Sprintf("    %s: count=%d, total=%v, avg=%v, min=%v, max=%v\n",
				cmd, m.Count, m.TotalTime, m.Avg(), m.MinTime, m.MaxTime))
		}
	}
	return sb.String()
}

func (p *ProfilerPlugin) Reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.byteMetrics = make(map[int]*Metric)
	p.stringMetrics = make(map[string]*Metric)
	p.totalByteCalls = 0
	p.totalStringCalls = 0
	p.totalByteTime = 0
	p.totalStringTime = 0
	p.startTime = time.Now()
}

func (p *ProfilerPlugin) Enable() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.enabled = true
}

func (p *ProfilerPlugin) Disable() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.enabled = false
}

func (p *ProfilerPlugin) Close() error {
	return nil
}

func (p *ProfilerPlugin) Call(name string, _ ...core.Option) (any, error) {
	switch name {
	case "report":
		return p.Report(), nil
	case "reset":
		p.Reset()
		return "reset done", nil
	case "enable":
		p.Enable()
		return "enabled", nil
	case "disable":
		p.Disable()
		return "disabled", nil
	default:
		return nil, fmt.Errorf("unknown call: %s", name)
	}
}

func (p *ProfilerPlugin) Run(any) (any, error) {
	return nil, nil
}
