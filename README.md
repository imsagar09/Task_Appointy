# Task_Appointy


### /meetings
  * method POST: To schedule a meeting, provide a Title(string), Participants(array of participant data in valid format), Start Time, End Time)
  eg - 
  <br>{
"Title": "Meeting",<br>
"Participants":[{"Name": "sagar", "Email": "xyz@gmail.com", "RSVP": "Yes"}, <br>{"Name": "gupta", "Email": "sag@gmail.com", "RSVP": "No"}],<br>
"Start Time":"2020-10-20T13:11:10Z",<br>
"End Time": "2020-10-20T13:14:00Z"
 }<br><br>
 
returns meeting_id
### /meetings?participant=mail_id_of_participant
  * method GET
   eg -<br>
   localhost:8082/meetings?participant=xyz@gmail.com
 
   returns array of meetings<br>
### /meetings?start=start_time&end=end_time
  * method GET
  eg - <br>
  localhost:8082/meetings?start=2020-10-20T12:13:00Z&end=2020-10-20T15:14:00Z
  <br>
  returns array of meetings
  
### /meeting/meeting_id
  * method GET
    eg -<br>
    localhost:8082/meeting/7qwa54sdr5vgh56ytg48ed
    <br><br>
   returns meeting
 
