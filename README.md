# Job scrapper
## Clone coding from https://nomadcoders.co/go-for-beginners/lobby (#4, #5)
- 2022년 여름방학 인턴 교육 과정의 일환으로 진행했으나 기록을 남기지 않았기도 했고 
- go echo 프레임워크에 대해 관심이 생겨서 review
## 스스로 수정한 점
원래 https://kr.indeed.com/ 에서 job scrap을 하는 내용이지만 현재 해당 사이트에 크롤링이 막혀서(status code 403)   
대신 https://www.saramin.co.kr/ 에서 수행하게 코드를 변경
- URL의 argument 형식이 달라서 이에 맞게 수정
- HTML의 class 이름과 id로 해당 구인 페이지를 찾는 방식이 달라 이에 맞게 수정
  - ex) indeed의 경우 jk=<id값> 으로 페이지를 찾고 saramin의 경우 rec_idx=<id값>으로 페이지를 찾음
