# Robonona

Robonona is a Mattermost bot that, through BambooHR's API, posts birthdays and anniversaries on any specified channel.
## Installation

1. Install Go
2. Create a directory for your go code (if necessary)
3. Set your `GOPATH` environment variable to point to your Go directory
4. Run `go get github.com/garcijo/robonona`

## Usage

1. Go to your `$GOPATH/src/github.com/garcijo/robonona` directory.
2. Create a file (we call ours `.env` that looks like the following:
```
ROBONONA_MATTERMOST_URL="https://<your mattermost server url>"

ROBONONA_USERNAME="<some mattermost username>"
ROBONONA_PASSWORD="<that username's password>"

ROBONONA_TEAM_NAME="<the name of some mattermost team name>"
ROBONONA_CHANNEL_NAME="<the name of some mattermost channel under that team>"

ROBONONA_BAMBOO_URL="<the URL for some company's BambooHR page>"
ROBONONA_BAMBOO_KEY="<the user's BambooHR API key>"
```
3. `go run main.go`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Future Features
- Compiled celebrations (multiple birthdays/anniversaries on the same day)
- Statutory holidays
- Custom holidays (bot's birthday, company's anniversary, etc.)
- Direct congratulatory messages

## License
This software is licensed under the [MIT](https://choosealicense.com/licenses/mit/) software license.
