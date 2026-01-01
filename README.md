# Gluttony (zzizily-gluttony-go)

**예약 자동화 도구**

`Gluttony`는 `go-rod`를 기반으로 한 브라우저 자동화 도구입니다. 설정된 시나리오에 따라 지정된 웹사이트에 로그인하고, 특정 예약을 자동으로 수행하거나 스크린샷을 캡처하는 기능을 제공합니다.

## 주요 기능

- **자동 로그인**: 환경 변수에 저장된 자격 증명을 사용하여 자동 로그인을 수행합니다.
- **시나리오 기반 실행**: YAML 설정 파일을 통해 수행할 작업(로그인, 예약 등)을 순차적으로 정의하고 실행할 수 있습니다.
- **예약 자동화**: 캘린더 등에서 특정 조건의 이벤트를 찾아 클릭하고 예약을 진행합니다.
- **스크린샷 캡처**: 각 단계별로 스크린샷을 찍어 `screenshot/` 폴더에 타임스탬프와 함께 저장합니다.
- **로깅**: 실행 로그를 `logs/YYMMDD/` 디렉토리에 날짜별로 기록합니다.

## 시작하기 (Getting Started)

### 1. 요구 사항 (Prerequisites)

- Go 1.18 이상
- Google Chrome (go-rod가 내부적으로 브라우저를 제어함)

### 2. 설치 및 설정

**레포지토리 클론**

```bash
git clone https://github.com/deuxksy/zzizily-gluttony-go.git
cd zzizily-gluttony-go
```

**환경 변수 설정**
`.env.example` 파일을 복사하여 `.env` 파일을 생성하고, 필요한 정보를 입력합니다.

```bash
cp .env.example .env
```

`.env` 파일 내용 편집:

```ini
USERID=your_username  # 로그인 ID
USERPW=your_password  # 로그인 패스워드
GO_PROFILE=local      # 사용할 설정 프로필 (기본값: local)
```

**시나리오 설정**
`configs/` 디렉토리에 설정 파일(예: `local.yml`)을 작성합니다.

```yaml
Scenario:
  - name: Login
    url: https://example.com/login
    type: login
  - name: Healthcare
    url: https://example.com/rb/main#booking/calendar?resourceId=...
    type: booking
```

- **Login**: `type: login`으로 설정하며, 로그인을 수행합니다.
- **Booking**: `type: booking`으로 설정하며, 해당 URL로 이동하여 예약을 시도합니다.

### 3. 실행 (Usage)

**직접 실행 (Go Run)**

```bash
go run cmd/gluttony/main.go
```

**빌드 후 실행**

```bash
# Windows
go build -o gluttony.exe cmd/gluttony/main.go
# Mac/Linux
go build -o gluttony cmd/gluttony/main.go

./gluttony
```

## 프로젝트 구조

```text
.
├── cmd/
│   └── gluttony/       # 메인 애플리케이션 엔트리 포인트
├── configs/            # 시나리오 설정 파일 (YAML)
├── internal/
│   ├── configuration/  # 설정 구조체 및 파싱 로직
│   └── logger/         # 로깅 설정 (Zap)
├── logs/               # 실행 로그 저장소 (자동 생성)
├── screenshot/         # 스크린샷 저장소 (자동 생성)
├── .env                # 환경 변수 (Credentials)
└── go.mod              # Go 모듈 의존성
```

## 출력 결과

- **스크린샷**: `screenshot/YYMMddHHmm/` 디렉토리에 시나리오 실행 단계별 이미지가 저장됩니다.
- **로그**: `logs/YYMMdd/` 디렉토리에 `out.log` (일반 로그)와 `error.log` (에러 로그)가 저장됩니다.
