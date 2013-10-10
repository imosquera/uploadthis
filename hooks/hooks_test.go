package hooks

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/util/mocks"
	. "launchpad.net/gocheck"
	"testing"
)

func TestHooks(t *testing.T) { TestingT(t) }

type HookSuite struct{}

var _ = Suite(&HookSuite{})

func (s *HookSuite) TestGetPrehooks(c *C) {

	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	prehooker := mocks.NewMockCommander(mockCtrl)
	mockPrehooks := []string{"mock_prehook"}
	prehookMap := make(map[string]commands.Commander, 0)
	RegisterHook(PREHOOK, "mock_prehook", prehooker)

	GetHookCommands(PREHOOK, mockPrehooks, prehookMap)
	_, ok := prehookMap["mock_prehook"]

	c.Assert(len(prehookMap), Equals, 1)
	c.Assert(ok, Equals, true)
}

func (s *HookSuite) TestRegisterCompressHook(c *C) {
	prehook := &CompressPrehook{}
	RegisterHook(PREHOOK, "mock_name", prehook)
	c.Assert(registeredPrehooks["mock_name"], Equals, prehook)
}
