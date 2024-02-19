package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("repeat 'hello' 3 times", func(t *testing.T) {
		got := Repeat("hello", 3)
		want := "hellohellohello"
		assertCorrectMessage(t, got, want)
	})

	t.Run("repeat 'a' 1 times", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"
		assertCorrectMessage(t, got, want)
	})

	t.Run("repeat 'a' 1 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
