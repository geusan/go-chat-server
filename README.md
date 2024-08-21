# Chat server

채팅서버 튜토리얼 작성

## 목표

- 채팅 서버를 생성한다.
- 채팅방의 최대 규모는 5인으로 제한한다.
- 채팅의 콘텐츠는 텍스트와 이미지로 제한한다.

## Libraries
- echo server
- ORM: gorm
- Docs: swagger
- testing: mockery

## 계획
- 채팅 기능은 소켓을 활용한다.(gorilla socket)
- 채팅 클라이언트 React로 작성
- [ ] 채팅 서버 개발
    - [ ] 채팅 서비스 API 서버, 채팅 서버 분리
- [x] API 서버 구축 with echo
- [x] swagger를 이용한 API document 툴 개발
- [x] ORM을 이용 (gorm)
- [x] unit test 작성 연습 포함(Mockery)

## 명령어

Swag
```bash
# swag 실행시 gorm.Model을 못찾는 이슈가 있는데 내부 디펜던시 관련 옵션 포함
swag init --parseDependency --parseInternal

```

Mockery
```bash
# 전체 경로 mocks 생성
go generate ./...
```

DB
```bash
# gorm의 AutoMigrate 활용
# TODO: migration manager에 대한 연구 필요
go run cli/migration.go
```

## 추가 작업
- [x] 기능이 변경되며 발생한 rdb, service, rest 등 테스트 추가 작업
- [x] middleware 테스트 추가
- [ ] DTO와 AddUser 등 이름이 다른 DTO 이름 통일
