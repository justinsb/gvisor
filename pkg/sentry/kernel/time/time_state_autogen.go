// automatically generated by stateify.

package time

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (t *Time) StateTypeName() string {
	return "pkg/sentry/kernel/time.Time"
}

func (t *Time) StateFields() []string {
	return []string{
		"ns",
	}
}

func (t *Time) beforeSave() {}

func (t *Time) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.ns)
}

func (t *Time) afterLoad() {}

func (t *Time) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.ns)
}

func (s *Setting) StateTypeName() string {
	return "pkg/sentry/kernel/time.Setting"
}

func (s *Setting) StateFields() []string {
	return []string{
		"Enabled",
		"Next",
		"Period",
	}
}

func (s *Setting) beforeSave() {}

func (s *Setting) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Enabled)
	stateSinkObject.Save(1, &s.Next)
	stateSinkObject.Save(2, &s.Period)
}

func (s *Setting) afterLoad() {}

func (s *Setting) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Enabled)
	stateSourceObject.Load(1, &s.Next)
	stateSourceObject.Load(2, &s.Period)
}

func (t *Timer) StateTypeName() string {
	return "pkg/sentry/kernel/time.Timer"
}

func (t *Timer) StateFields() []string {
	return []string{
		"clock",
		"listener",
		"setting",
		"paused",
	}
}

func (t *Timer) beforeSave() {}

func (t *Timer) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.clock)
	stateSinkObject.Save(1, &t.listener)
	stateSinkObject.Save(2, &t.setting)
	stateSinkObject.Save(3, &t.paused)
}

func (t *Timer) afterLoad() {}

func (t *Timer) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.clock)
	stateSourceObject.Load(1, &t.listener)
	stateSourceObject.Load(2, &t.setting)
	stateSourceObject.Load(3, &t.paused)
}

func init() {
	state.Register((*Time)(nil))
	state.Register((*Setting)(nil))
	state.Register((*Timer)(nil))
}
