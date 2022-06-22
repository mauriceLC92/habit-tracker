## Questions for session
- why is the test file it's own package? Is that so that we can import our package into it?
    - The answer to this is that we are treating the tests as a client of our package. It needs to follow the same rules
        that a real user of our package would need to follow. Ask yourself why should it be any different?

- for an application where our engine package is the habit package, how does one add on things for web vs the cli?
For example, when creating a new habit, we will log that a new one has been created to the logger and not stdout

## Notes for next session
- I struggled quite a bit to get started with a test first approach.
- When writing a test for retrieving a habit by ID it was weird for me to think where should I be fetching this habit from?
    - For example) Does one write the function and able to do so in a way that the user can provide their own persistance layer? How does one do that?
    - So is a user able to either provide that they want to write to a database or to the file system - How is that accomplished?

- Should a folder always be a package? How should my db folder be structured?