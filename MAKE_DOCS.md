### Basic Api with Health Check With Docker Compose

#### Make Commands

Starting point for these helpful commands is your Makefile.
Most of the commands in the makefile can easily be ran independently.

Each of the following commands should be prefixed with `make`. Example: for `build`, run `make build`.

`build-deps` : 
    Ensures your "backend" service has all go dependencies specified in your lock file.
	
`build`: 
     Re-build/build the "backend" service as a docker image.

`build-with-deps`: 
    should be ran only the first time building the "backend" service 

`run`:
    Default Make command, alias of `up`

`start`:
    Alias of `up`

`up`:
	Starts the docker compose daemon.

`stop`:
    Alias for `down`
    
`down`:
	Bring docker-compose images down.

`restart`:
	Restart docker-compose

`log`: 
    Alias of `logs`
    
`logs`:
	Show docker-compose logs
    
`logf`: 
   Follow docker-compose logs

`env`:
	Create .env file from example

`envs`:
	Print current environment variables used by the "backend" service

`enter`:
	Run a prompt inside the "backend" service's docker image

`test`:
	Run tests inside the "backend" service's docker image

`test-local`:
	run tests locally

### Dependencies

At least:
    Docker v1.9.0
    Docker Compose v1.5.0
    Docker Machine v0.5.0
    make   

### Credits

Fork of [Go Docker](https://github.com/typenil/go-docker)
