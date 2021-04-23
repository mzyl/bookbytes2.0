## BookBytes 2.0
#### Short book passages to pique your interest
Things are coming together.
Interface is HEAVILY borrowed from the clickbaiter project by cbrgm found here: https://github.com/cbrgm/clickbaiter/

### Goals
- [x] Better data passing through functions
- [x] New buttons for displaying title and author
- [x] Build "booktext" out of "fulltext"
- [x] Maybe begin random selection of paragraphs and/or books
- [x] Random file selection
- [x] Select new book from ui
- [ ] Address some "problems"
- [x] Better layout - Buttons on the side, maybe
- [x] Next paragraph button
- [x] Previous paragraph button
- [x] Beginning of chapter button
- [x] Display title/author and the end at beginning and end of book
  - Maybe only at the end? 
  - Append "The End" to final chapter
    - Final chapter ends with *Fin.*

### Goals for Release
- [x] Able to display title and author
- [x] Random selection of books and paragraph
- [x] Basic navigation through books i.e. forward, backward, beginning of chapter
- [x] Solution for reading more of the book i.e. forward-backward are not so great for dialog
  - Field is now scrollable
- [x] Work with whole book library
- [ ] Host locally on server
- [ ] Find font that contains virtually all characters

### Problems
- Chapters are not all marked the same.....
  - This is working better, but not perfectly.
- Same book is served to all users and changes made, i.e. new book/paragaph, effect all users.
- Scrollable page and div doesn't work well on mobile.
  - Also seems to be a slight input lag on mobile?
- 3/1/3/4/31342/31342-h/313420h.htm returns only "Fin."
  - it is a small collection of poems.
  - encapsulated by "table" instead of "p".
- Poe poems are encapsulated in "pre . . ." so they aren't pulled with "p"'s.

### Thoughts
- Eventually would like to have accounts to save the reader's place.
- May want to filter for English texts.
- Selector file runs separately when server starts.
  - Hold current possible files that can be served using a different file.
    - May have been taken care of with File variable addition?

### Ordered Plan
- [x] Set up server to download full library
- [x] Pull html documents out of full library
  - rsync can do this for me
- [ ] Explore ideas for segmenting chapters and other print section formats
  - Lots of progress made here thanks to PyrO
- [ ] Host locally from server
- [ ] Reimplement so each connection to site gets a different session
  - ~~Also should store library files to be accessed repeatedly instead of walking directory every time~~
    - "Files" are now stored in global variable
- [ ] Add feedback solution
- [ ] Finalize web interface
- [ ] Replicate local server on digitalocean or similar
- [ ] Link to domain name
- [ ] Release and look for feedback
