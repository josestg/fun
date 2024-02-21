package fun

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunWithContext(t *testing.T) {
	dummyErr := errors.New("some error")

	timed := func(delay time.Duration) func() (int, error) {
		return func() (int, error) {
			time.Sleep(delay)
			return 42, dummyErr
		}
	}

	t.Run("no cancel", func(t *testing.T) {
		ctx := context.Background()
		got, err := RunWithContext(ctx, timed(time.Millisecond))
		if !errors.Is(err, dummyErr) {
			t.Errorf("expect %v; got %v", dummyErr, err)
		}
		if got != 42 {
			t.Errorf("expect 42; got %d", got)
		}
	})

	t.Run("deadline exceeded", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()

		got, err := RunWithContext(ctx, timed(time.Second))
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("expect %v; got %v", context.DeadlineExceeded, err)
		}
		if got != 0 {
			t.Errorf("expect 0; got %d", got)
		}
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(100*time.Millisecond, cancel)
		got, err := RunWithContext(ctx, timed(time.Second))
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expect %v; got %v", context.DeadlineExceeded, err)
		}
		if got != 0 {
			t.Errorf("expect 0; got %d", got)
		}
	})

	t.Run("cancel before start", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel before starting the function.

		got, err := RunWithContext(ctx, timed(time.Second))
		if !errors.Is(err, context.Canceled) {
			t.Errorf("expect %v; got %v", context.DeadlineExceeded, err)
		}
		if got != 0 {
			t.Errorf("expect 0; got %d", got)
		}
	})
}

func TestRunWithTimeout(t *testing.T) {
	dummyErr := errors.New("some error")

	timed := func(delay time.Duration) func() (int, error) {
		return func() (int, error) {
			time.Sleep(delay)
			return 42, dummyErr
		}
	}

	t.Run("no cancel", func(t *testing.T) {
		got, err := RunWithTimeout(time.Second, timed(time.Millisecond))
		if !errors.Is(err, dummyErr) {
			t.Errorf("expect %v; got %v", dummyErr, err)
		}
		if got != 42 {
			t.Errorf("expect 42; got %d", got)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		got, err := RunWithTimeout(time.Millisecond, timed(time.Second))
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("expect %v; got %v", context.DeadlineExceeded, err)
		}
		if got != 0 {
			t.Errorf("expect 0; got %d", got)
		}
	})
}

func TestRunWithDeadline(t *testing.T) {
	dummyErr := errors.New("some error")

	timed := func(delay time.Duration) func() (int, error) {
		return func() (int, error) {
			time.Sleep(delay)
			return 42, dummyErr
		}
	}

	t.Run("no cancel", func(t *testing.T) {
		got, err := RunWithDeadline(time.Now().Add(time.Second), timed(time.Millisecond))
		if !errors.Is(err, dummyErr) {
			t.Errorf("expect %v; got %v", dummyErr, err)
		}
		if got != 42 {
			t.Errorf("expect 42; got %d", got)
		}
	})

	t.Run("deadline exceeded", func(t *testing.T) {
		got, err := RunWithDeadline(time.Now().Add(time.Millisecond), timed(time.Second))
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("expect %v; got %v", context.DeadlineExceeded, err)
		}
		if got != 0 {
			t.Errorf("expect 0; got %d", got)
		}
	})
}
