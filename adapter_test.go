package adapter

import "testing"

func TestAdapter(t *testing.T) {

	keyBoard := NewKeyBoard()
	adapter := NewAdapter(keyBoard)

	t.Log(keyBoard.UseUSB())
	t.Log(adapter.UseTypeC())
}
