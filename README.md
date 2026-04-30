# MagisterLinguae 0.4.1
A tool that helps u learn languages

The name comes from latin, Magister linguae is Latin for master of language or teacher of language. I thought it would be fitting for this kind of project
The project is my attempt at creating language learning engine fitting for my needs. Any features added will fit my use case.
It is intented to be my solely my project, partly for my CV and partly for my use. But feel free to fork it or use it as u wish.
No AI was used for generating this code, it might be slop, but it's my slop. AI was used only for help with direction and what to focus on. And occasionly
the sneaky bugs and for some ideas how to refactor my code. No code generated, or altered by AI:)

MagisterLinguae is based on the input theory learn more here: https://en.wikipedia.org/wiki/Input_hypothesis MagisterLinguae is supposed to help u in this process
by marking which words u know, how often u encounter these words and single out the most frequent unknown ones if needed and also refresh the ones u have not seen a lot time.
The more u use it, the more useful the app will be. The goal is to get you to the point u understand most of the commonly used language, the grade reading goal is at least B1. 
Then you can start learning grammar and practice your output, but not here, there is no grammar or no place for your output. There will be an option for dictionary translation,
but only for single words or short phrases, you're not really supposed to translate whole sentences. 

The project is written in Go, as it's my most used language, It's currently REPL, web UI will be coming later.

This is the alpha version, anything included is prone to bugs and partly correct implementations, it should be believed that everything in this build will probably
not stay the same in the release of 1.0

Current version is 0.4.1, I tried to add more tests, but fell short. I would have to simulate user input in a lot of functions which would mean a lot of rewriting, I'm just gonna finish
the app, then test everything

## Current features: 
- The program can recognise words (case insensitive) which are then stored in db
- Words are stored stored with boolean known and unknown, user marks words that he knows, any other are unknown*
- Words are arranged back into previous sentences, unknown words are distinguishable by these brackets "[]" in the stdout
- Currently supports only Italian, other languages haven't been tested

*known words should be marked when user completely understands them 

These all will be updated depending on what is completed and what new comes up
### Future features (dopamine corner):
(based on priority and its complexity, most important and least complex are on top)
- refactorization of the code, make it simpler, cleaner and more readable ✅
- lookup, u can search word words in db, if they are there or not and the last time u seen them + frequency ✅
- probably REPL, makes the use of program very easier, at least for marking ✅
- prompting the user to mark words as known after reaching certain frequency of the word ✅
- tests ✅
- user manually changing the words he forgot to unknown ✅
- ranking the words based on frequency, especially the unknown ones, so you know which one to learn ✅
- add support for files ✅
- help command ✅
- subtitles and youtube transcript compatibility ✅
- supporting other languages ✅
- ability to create multiple own instances of languages/tabs ✅
- dictionary api ✅
- quiz based on words u know and have not seen for a long time, kinda like anki or something like that ✅
- L'italiano Secondo Il Metodo Natura transcript available as a start. (maybe even more graded books) ✅
- docker support ❌ decided against it as it adds low benefit, will be added once the project has web UI

everything above will be in the 1.0 version
(The holy grail)
- web UI - version 2.0

#### Optional:
- AI support
- account support
- support for languages with different writing
- additional resources for input learning
- text to speech
- subscription or paygate system (just for learning, I intend for this project to stay open source)

## Instalation guide:
### Prerequisites

- **Go** 1.24.3 or later
- **PostgreSQL** 12 or later
- **Goose**

### 1. Clone the Repository

```bash
git clone https://github.com/DXICIDE/MagisterLinguae
cd MagisterLinguae
```

### 2. Create the Databases
```bash
sudo -u postgres psql   # Linux/macOS
psql -U postgres         # Windows
```

#### inside create these db
```sql
CREATE DATABASE words;
CREATE DATABASE magisterlinguae_test;
\q
```

### 3. Configure environment
Create a .env file in the project root:
```go
DB_URL="postgres://postgres:YOUR_PASSWORD@localhost:5432/words"
TEST_DB_URL="postgres://postgres:YOUR_PASSWORD@localhost:5432/magisterlinguae_test"
```


### 4. Install goose if u dont have and run the migrations
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir sql/schema postgres "postgres://postgres:YOUR_PASSWORD@localhost:5432/words" up
```
#### ⚠️ Migration 005_newprimkey.sql clears existing word data (TRUNCATE). This only affects you if you're re-running migrations on an existing database. ⚠️

### 5. Final
```bash
go run .
```
#### You should see:
```bash
Select language to learn please:
1 - it - Italian
2 - cz - Czech
Select language by typing its number:
```
### Troubleshooting:

- **connection refused**  -   Ensure PostgreSQL is running: sudo systemctl start postgresql
- **password authentication**  -  failed Check that the password in .env matches your PostgreSQL password
- **database "words" does not exist** -	Run Step 2 again
- **goose: command not found** -	Run Step 4 again; ensure $GOPATH/bin is in your PATH
- **Tables don't exist or errors on startup** -   Re-run migrations from Step 5