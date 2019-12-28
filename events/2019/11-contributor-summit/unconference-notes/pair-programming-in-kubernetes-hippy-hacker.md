#### Presenter: Hippy Hacker

#### Topic: Pair Programming in Kubernetes using EMACS

#### Date & Time: 3:30 PM

#### Notes

1. So you will get ssh link
2. You ssh both people ssh
3. There can be a todo list 
4. When you check todo list off it will show percentage updated
Basically a tmux session uses emac

Basically opens a tmux interactive window session where you can execute code directly from the shared document
You can do a left and right `i`, this will allow you to control things from one left `i` and execute on the right `i` in 
real time. Both eyes are accessed via links
The session is shared so both people can pair.

The desired use is to to write test in a paired fashion.
When ou write test you can leverage kind to start clusters for your tests to be able to identify which api's your using.

There are multiple helper functions included like `Todo` `Blocked` `Done` that allow the pair programmers to track 
status and also you can get your test results directly in the document while peer programming.

In the right `i` you can actually run through logs and show what happened with your tests.


The [repo](github.com/linmacs/emacs.d) has the instructions to install and configure the pair programming.


CNCF - contract is online to help with conformance working group to write api stint to be able to deploy into cluster 
and see what is happening in the tests.

### Take Aways
need to help community adoption along

### Action Items
* Get SIG leaders together to help with adoption and understand who to reach out to when tests aren't created?
