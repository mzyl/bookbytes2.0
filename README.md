## BookBytes 2.0
#### Short book passages to pique your interest
Things are coming together.

### Goals
- [ ] Address some "problems"
- [x] Begin addressing the sessions problem
    - Maybe try implementing sessions elsewhere, then bring the idea back here
- [x] Text field reset on button press
- [x] Grab and display language in Info
    - Will be useful for filter option later
- [ ] Clean up/Refactor functions

### Goals for Release
- [x] Able to display title and author
- [x] Random selection of books and paragraph
- [x] Basic navigation through books i.e. forward, backward, beginning of chapter
- [x] Solution for reading more of the book i.e. forward-backward are not so great for dialog
    - Field is now scrollable
- [x] Work with whole book library
- [ ] Host locally on server
- [ ] Find font that contains virtually all characters
    - Maybe this doesn't exist like I thought it did
- [ ] About page, what else is coming, etc.

### Ordered Plan
- [x] Set up server to download full library
- [x] Pull html documents out of full library
    - rsync can do this for me
- [ ] Explore ideas for segmenting chapters and other print section formats
    - Lots of progress made here thanks to PyrO
        - Needs more refining before release
- [ ] Host locally from server
- [x] Reimplement so each connection to site gets a different session
- [ ] Finalize web interface
- [ ] Add feedback solution
- [ ] Replicate local server on digitalocean or similar
- [ ] Link to domain name
- [ ] Release and look for feedback

### Problems
- Chapters are not all marked the same.....
    - This is working better, but not perfectly.
- Same book is served to all users and changes made, i.e. new book/paragaph, effect all users.
    - This was fixed by passing filename data to and from browser on each request.
- Scrollable page and div doesn't work well on mobile.
    - Also seems to be a slight input lag on mobile?
- 3/1/3/4/31342/31342-h/313420h.htm returns only "Fin."
    - it is a small collection of poems.
    - encapsulated by "table" instead of "p".
- Poe poems are encapsulated in "pre . . ." so they aren't pulled with "p"'s.

### Thoughts
- Some paragraphs are a little daunting.
    - Maybe we should implement a character limit as well?
- Should consider checking for sufficient content before displaying selected book.
    - If we can't figure out chapter/paragraph issues for all texts,
    - maybe we just display the texts we have working sufficiently for release.
- ~~Maybe store list of files in text for Go to reference instead of rescanning the files on startup.~~
    - This has been implemented with the help of ikiris.

### Things that could be improved
- [ ] Chapter selection could use some help

### Future ideas
- Eventually would like to have accounts to save the reader's place.
    - Accounts may be able to generate random paragraphs from specific books they have enjoyed in the past.
- May want to filter for English texts.

### Issues found with text parsing
#### Need to replace most Regexp with Tokenizer
- [ ] Section wrapped in pre or span, not p.
- [x] Licensing information wrapped in p and span, not plain text.
- [x] No START or END, no text is appended in StripLiecense().
    - StripLicense stores text in fulltext which is used for booktext.
        - Meaning there is no text in booktext.
- [x] No spaces between \*\*\* and Start/End.
    - Also "Start/End of THE", not THIS.
- [x] 27210 splits lines on "Start"
    - Do we want to handle new lines or just remove the end \*\*\*?
- [ ] Some html files just hold space for mp3 files and not books
    - Maybe check for things like this before serving the book?
    - These files may not have a Start and End like the others.
- [ ] 16446 is a set of poems, no p tags anywhere
- [ ] ../library/htmlmirror/5/6/6/5669/old/2004-05-conrg10h.htm
    - No ending p tags
