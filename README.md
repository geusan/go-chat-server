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
- [x] 채팅 서버 개발
    - [x] 채팅 서비스 API 서버, 채팅 서버 분리
- [x] API 서버 구축 with echo
- [x] swagger를 이용한 API document 툴 개발
- [x] ORM을 이용 (gorm)
- [x] unit test 작성 연습 포함(Mockery)

## 계획2
- [ ] 리팩토링 계획
    - [ ] 프론트엔드 타입 정리
    - [ ] 백엔드 서버 에러 코드 정리
- [ ] 채팅 서버 관리(안정해시) 저장소 변경(in-memory => redis cluster)
    - [ ] 채팅서버 실행시 Redis에 저장상태 on 하도록 수정
    - [ ] 채팅 서버에 문제가 있다,없다는 것을 어떻게 관리? 고민(스케쥴잡 ping 서버를 만들고 시계열로 전송?)
- [ ] 각 채팅방에 현재 몇명이 있는 지 확인
- [ ] 프론트내에 현재 접속중인 사용자 명수 표기(polling 방식으로 구현)
- [ ] N개의 채팅서버와 1개의 API 서버, RDB, Redis를 관리하는 테라폼 작성
    - [ ] 개발 편의성상 코드 변경시 컨테이너가 재실행되도록 환경 구현
- [ ] 빌드 프로세스를 Makefile로 작성

## 계획3
- [ ] 프록시 서버를 만들고 로컬에서 https 구현(개발 편의상..)
    - [ ] https 한 김에 web-push도 구현
- gRPC 도입
- webRTC 실험(영상, 음성 통화 가능성 확인)


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
