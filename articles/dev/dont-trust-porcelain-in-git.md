---- Jul 14 13:37:41 IST 2023
# Don't trust porcelain in Git

## Authors of Git decided to use toilet analogy when trying to differentiate the low-level stuff from the high-level stuff. The analogy works well in all but one case.

Git commands can be classified in many different ways but here is one:
1) There are git commands whose output is meant __to be read by humans__, for example "git status". They are everyday commands that "normal" developers use to version their projects. You can call these commands "high-level" because they are built on top of commands from the second group. 
2) There are git commands (such as "git ls-files") whose output is meant __to be consumed by external scripts__. As git documentation puts it, "Many of these commands aren’t meant to be used manually on the command line, but rather to be used as building blocks for new tools and custom scripts".

To clarify this distinction, Authors of Git suggested this slightly filthy analogy. High-level stuff they called "porcelain", low-level - "plumbing". These words must help you quickly understand the idea by evoking the images of a restroom in your mind. Because toilets are made of porcelain, right?

This is how I understand it:


user1, user2, toilet, pipes.

User1 is probably you, a common git user. You might be rightfully offended by this metaphor. But I'm not to blame, I'm just following Fathers' logic.
As a User1 you only deal with high-level commands, you just read their output using your eyes, and that's it.

But for machines this "porcelain" output is less digestible and, what's more, this output is not guaranteed to be stable, i.e. Git can change the format of these messages in its later versions without notice.

So if you are a User2 (the one who uses git output in his own scripts) then you should't rely on porcelain output. Instead you should go for plumbing which provides stable and machine-readable information.
For example your script needs a kind of data that is provided by "git status". But you shouldn't do it because git status is a porcelain command. Its output 1) will perhaps be harder to parse by your script and 2) can be changed in the future, therefore breaking your script. (This is by the way the only practical advice that you can take from this article).

You need a "plumbing" command in such case. There must be some very clever command from git internals for that (I'm not aware of such command, just rambling). But Git also offers a more user-friendly way - you can run "git status" with a "--porcelain" flag and get a machine-readable data in a stable format, in other words... a "plumbing" data.

Then why is this flag called "--porcelain" at all? Who knows. Many fussy developers think that this name is mistaken. But the flag does the job so you can just memorize it and be happy.

(And of course this story with the flag does not undermine the overall concept of porcelain vs plumbing).
