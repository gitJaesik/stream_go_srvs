# stream_go_svrs

## 포트폴리오를 위한 소스 작업 중

### streamgolib

- grpc server들이 사용할 공통 라이브러리 및 wrapper 함수들

### stream-auth-server-go

- jwt 인증 및 iam 정보를 담당할 서버

### stream-file-server-go

- 스트림에 대해 실시간 파일 저장 및 unary 파일 저장등을 담당할 서버
- s3와 비슷하지만 다른 것은 중간 패킷 저장 가능 (클라이언트 메모리 부하 줄임)

### stream-api-server-go

- 메인 로직을 담당할 서버
