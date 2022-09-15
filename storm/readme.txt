1. mysql doker 

docker run --platform linux/amd64 --name mysql-storm -e MYSQL_ROOT_PASSWORD=storm -d -p 3306:3306 mysql:5.7.39

macos m1 에서 mysql 사용
https://www.lainyzine.com/ko/article/how-to-install-docker-for-m1-apple-silicon/

docker cp /tmp/테이블명.sql 컨테이너ID:/tmp  	 	 	# 컨테이너에 파일 복사하기 

docker exec -it 컨테이너ID sh 			  	 	# 컨테이너에 접속해서 

DB 생성

CREATE DATABASE storm;

ALTER DATABASE
    storm
    CHARACTER SET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;


기존 데이터 넣기.
import 시에 한글 있는 경우 케릭터셋을 지정해서 넣어야한다.
mysql -p storm --default-character-set=utf8mb4 < xe_member_attend_check.sql

2. project 환경

fiber + mysql + gorm

mkdir storm
cd storm
go mod init storm
mkdir service models storage

go get -u github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10


SELECT CCSA.character_set_name FROM information_schema.`TABLES` T,
       information_schema.`COLLATION_CHARACTER_SET_APPLICABILITY` CCSA
WHERE CCSA.collation_name = T.table_collation
  AND T.table_schema = "storm"
  AND T.table_name = "xe_member_attend_check";


SELECT COLUMN_NAME, character_set_name,COLLATION_NAME FROM information_schema.`COLUMNS` C
WHERE table_schema = "storm"
  AND table_name = "xe_member_attend_check"
  AND column_name = "m_name";  
