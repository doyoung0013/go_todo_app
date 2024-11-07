package config

import (
	"fmt"
	"testing"
)

// New 함수의 동작을 검증하는 테스트 함수
func TestNew(t *testing.T) {
	// Port 환경 변수 테스트를 위한 설정값
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	// Config 구조체 생성
	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	// Port 값이 환경 변수에서 정상적으로 읽혔는지 확인
	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}

	// Env 값이 기본값으로 정상 설정되었는지 확인
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}
