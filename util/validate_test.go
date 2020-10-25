package util

import "testing"

func Test_DeletePreAndSufSpace(t *testing.T) {
	if ans := DeletePreAndSufSpace("   asdsad   "); ans != "asdsad" {
		t.Error("DeletePreAndSufSpace测试失败-1")
	}

	if ans := DeletePreAndSufSpace("      "); ans != "" {
		t.Error("DeletePreAndSufSpace测试失败-2")
	}

	if ans := DeletePreAndSufSpace("  as as    "); ans != "as as" {
		t.Error("DeletePreAndSufSpace测试失败-3")
	}

	if ans := DeletePreAndSufSpace(""); ans != "" {
		t.Error("DeletePreAndSufSpace测试失败-4")
	}
}

func Test_AllIsSpace(t *testing.T) {
	t.Error(AllIsInvisibleCharacter("  s  "))
}
