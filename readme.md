# Youtube Sharing App
This project having 2 repos:

FE: https://github.com/votanlean/youshare

BE: https://github.com/votanlean/youshare-api

## Live Demo
Demo Link: https://youshare.vercel.app/

Demo Video: https://www.loom.com/share/7487e9cff9894cf59b9ba9f505e7945a?sid=92b1bfa1-e384-4aef-958c-04626a52ec6d

## Main features
- [ x ] Login
- [ x ] Register
- [ x ] View all videos
- [ x ] Share a youtube link
- [ x ] Get notification from newly added video via websocket

### Todo:
- [ ] User see new video instead of just new notification
- [ ] Unit test frontend
- [ ] Cover test backend
- [ ] Handle TLS versions for websocket connection
- [ ] Cache video
- [ ] Use queue to notify users asynchronously when heavy load on server


## Development - Frontend
(Repo: https://github.com/votanlean/youshare)

First, run the development server:

```bash
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

Environment Variables
```bash
NEXT_PUBLIC_API_URL=http://localhost:4000 # golang api
NEXT_PUBLIC_SOCKET_URL=ws://localhost:4000 # websocket api
NEXT_PUBLIC_GOOGLE_API_KEY=GOOGLE_API_KEY #for youtube api
```
You can change them or use the default one in the .env file
If you want to use .env in your localhost, clone it to .env.local

## Development - Backend
(Repo: https://github.com/votanlean/youshare-api)
Language: Golang.
I almost develop the CRUD and Websocket by builtin Golang function and just a few lightweight libraries

I've have add both unittest and integration test with testing database setup. However due to the time constraint I cannot cover the test coverage. So I just add the minimun test to showcase the test setup.

To run the project:
If you use makefile, you can run
```bash
make up # run the docker to get Postgres and create seed database
make down # remove docker and database
make run # run the project
```
Other wise you can directly run:
```bash
docker-compose -f deployments/docker-compose.dev.yml up -d
docker-compose -f deployments/docker-compose.dev.yml up
go run ./cmd/api # run the project
go build ./cmd/api # build the binary file
```
The localhost server is at [http://localhost:4000](http://localhost:4000)


## Deployment
The frontend is deployed to Vercel at https://youshare.vercel.app/


The backend is deployed to Digital Ocean and at domain https://youshare-api.anvo.dev/

Note that the backend is not stable because I have some troubles dealing with the postgres root user escalation bug when using Systemctl for detach process, as well as TLS for websocket 

In the localhost it works fine (as seen in demo video)







