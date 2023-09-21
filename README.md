# otp_based_authentication_2FA
Creating a 2FA OTP bases authentication backend framework in Golang


# flow

. Unhapppy path
1. invalid details provided
2. email already existed
3. number of attempts
4. otp not received
5. wrong otp

. Happy path

1. signup using email/phone and password
2. login using email/phone and password
3. send the otp.
4. validate the otp.

. API Signature:
- URL
- Request Method
- Request Headers
- Request Body
- Response Status code
- Response Headers
- Response Body

. Signup API:

- URL - https://random.com/api/signup

- Method - POST

- Content Type - application/json

- JSON body -
  {
    "email": "foo@bar.com",
    "password": "supersecret",
    "phone": "+91-9876543210"
  }

Validate request body
Insert record in DB
Handle errors
- Response Status code - 2xx, 3xx, 4xx, 5xx

. Login API

- URL - https://random.com/api/login

- Request method - POST

- Request body -
  {
    "email": "foo@bar.com",
    "password": "supersecret", // plaintext password, yikes!
  }
Logic on server side?

Validate request body
Check email and password combination
Handle errors
Create a token and send token back to user.
Response Status code: 2xx, 3xx, 4xx, 5xx

200 ok
5xx server side errors (DB down so not able to login the user)
Response Headers:

Content Type if we are sending any content back
Any other headers?
Return a time limited token in headers
Why do we need this? Let's come back to it when we talk about validate OTP API
Response Body:

Empty body or
Send OTP as part of response?
When do we actually send the OTP to user?

.Send OTP API

- URL https://random.com/api/v1/send-otp

- Request method: POST

- Request Headers:

- Token received from Login API reponse.

Request Body:

Empty body or
Should we send phone number to send OTP to?

Validate token from headers
Figure out phone number from token
Sent OTP, store OTP in some storage for validation
Response Status code: 2xx, 3xx, 4xx, 5xx

200 ok
4xx client side errors (invalid token, i.e. unauthorized)
5xx server side errors (Third party SMS vendor down, DB down, etc. so not able to send OTP the user)
Response Headers:

None that I can think of
Response Body:

Empty body


.Validate OTP API:

- URL https://random.com/api/validate-otp

- Request method: POST, or GET. Any query params?

Request Headers:

Token received from Login API reponse. Why?
any others?
Request Body:

JSON body
 {
    "otp": "123456"
 }
Logic on server side?

Validate token
Check token and OTP combination are correct
Handle errors
Response Status code: 2xx, 3xx, 4xx, 5xx

200 ok
4xx client side errors (invalid token, i.e. unauthorized, invalid otp, expired otp, etc.)
5xx server side errors (DB down so not able to validate OTP)
Response Headers:

None that I can think of
Response Body:

Empty body
Do you see any challenges/unknowns?
Token before and after OTP validation is same.
What if user requests for OTP multiple times? (both legit and abuser use case)
Where do we store OTPs, Tokens?
