# Changelog

이 프로젝트의 모든 주요 변경 사항은 이 파일에 문서화됩니다.

이 형식은 [Keep a Changelog](https://keepachangelog.com/ko/1.0.0/)를 기반으로 하며, 이 프로젝트는 [Semantic Versioning](https://semver.org/spec/v2.0.0.html)을 준수합니다.

## [Unreleased]

## [0.1.0] - 2026-01-02

### Added
- 프로젝트 변경 이력을 추적하기 위한 `CHANGELOG.md` 추가.
- `.gitignore`에 Go 빌드 및 OS 임시 파일 무시 규칙 추가.
- GitHub Actions를 이용한 자동 릴리스 설정 (`.goreleaser.yaml`, `.github/workflows/release.yml`) 추가.

### Changed
- `README.md` 전면 개편:
  - 프로젝트 개요, 설치/설정 방법, 사용법 상세화.
  - 디렉토리 구조 설명 추가.

## [2025-11-30] - Refactor

### Changed
- **Refactoring**: 하드코딩된 로직을 제거하고 설정 기반(Config-driven)으로 동작하도록 구조 변경. (by google-labs-jules[bot])
- Gitpod 환경 지원 추가.

## [2022-07 ~ 2022-10] - Initial Development

### Features
- **기능 안정화**:
  - `mustwait`, `waitload` 등을 적용하여 페이지 로딩 대기 로직 강화.
  - 로그인 후 Alert 창 무시 처리 (`ignore alert`).
  - 예약 로직 구현 ('헬스케어', '점심 식사' 등).
  - 로그인 모듈 분리 및 공통화.
- **환경 지원**:
  - OS별(Windows, Mac) 환경 변수 및 실행 스위치 처리.
  - Chrome Remote Debugging 및 GUI 모드 지원 추가.
- **스크린샷**:
  - 단계별 스크린샷 캡처 기능 구현.
  - 스크린샷 저장 폴더를 날짜/시간별로 자동 생성하도록 변경.
- **기타**:
  - `go-rod` 라이브러리 도입 및 테스트.

### Configuration
- `.gitignore`에 실행 파일 및 로그 폴더 추가.
- 스크린샷 및 로그 폴더명 충돌 방지 처리.

### Documentation
- Chrome 브라우저 설치 필요성 명시.
