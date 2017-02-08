version: '2'
services:
  mock:
    image: wanghaibo/httpmock
  test:
    image: *** 
    links:
      - mock:www.baidu.com
    
