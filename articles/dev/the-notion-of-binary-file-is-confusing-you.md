---- Aug 14 16:11:22 +07 2023
# The notion of "binary file" is confusing. You should stop using it

## My naive question is: how is it possible to talk about "binary" files at all? And does it make any sense to talk about "binary" files?

TLDR
All files are binary
Saying that file is a "binary file" is a kind of tautology similar to saying "digital file".
When you want to say "binary file", you probably mean "a file that doesn't contain human-readable text".
To make your speech more understandable, replace the notion of "binary" with "textual"/"non-textual".

On one hand, it's just a useless linguistic exercise.
On the other hand, it's very important. A notion of "binary" file seems to be very often misused. It pops up in various conversation and in most cases it either a) makes it harder to understand stuff or b) leads to incorrect understanding of stuff.

I wonder what is a non-binary file? Is it possible?
All files are binary deep inside, aren't they?
That is, somewhere deep inside they all consist of 0s and 1s.
And then this raw bits are interpreted and we get and image or text or a computational process.

I think it will be much cleaner to agree that all files are binary. And stop saying "binary file" as if there were some non-binary files.

So far there is a whole body of knowledge (or I should say pseudo-knoledge) in the Internet regarding binary files.
Though it seems that when people say "binary file", they actually mean something completely unrelated to binarity.

The most popular concept is that binary files are __non-textual__ files. Today (in 2023) an article on Wikipedia says exactly this:
""
This concept of __binary vs text files__ is indeed a convention in software industry. Everyone kind of agree on that but it doesn't make this concept less nonsensical. It's a serious naming issue at the core of many discussions.
Because at the same time nobody will argue that text files are binary too, right?
But somehow there is this assumption that text files are not "truly binary" because eventually they represent text.
And images for example are somehow __more binary__ than text.
Images qualify as "binary data" whereas text doesn't.
__Does it make any sense?__
I think it doesn't.
Well, ok, text files have certain features that distinguish them from other binary files, first of all that text files are very simple in structure. But it doesn't make them less binary.

I think we should stop saying "binary" to denote non-textual files. It would be much cleaner to talk about __human-readable and human-unreadable files__. Or "text files" and "non-text files". Binarity is simply of no relevance here.


Sometimes we talk about __compiling__ a source code into.... "binary"! Well, again it makes no sense. We can only compile one binary into another binary. What happens is that we compile a human-readable file to a human-unreadable file (and more efficiently executable by machine), alright? This explains stuff much better.
But... isn't machine code truly more binary than all other stuff? Maybe. Because machine code is "nothing more than binary"; it doesn't represent any "fancier" stuff like images or human-readable text. In case of machine code 0s and 1s are not translated into something else but executed directly by processor or CPU.
But even if we agree that machine code is "more binary", calling such files "binary" will be no less misleading. Perhaps "purely binary" should be used instead. Or, to avoid new terms, let's just call this stuff "machine code". Because it's exactly what it is.

And perhaps smart computer guys mean exactly that when they say "binary file". Or just "binary". They mean that it's not just binary but it's purely binary.
But what about stupid ordinary people? They take the notion of "binary files" at face value.

But what are __.bin__ files?
From my understanding, bin files are just... any files! Calling your file .bin doesn't make it behave in a certain way. And you can call any file (be it an image or text) ".bin" and I guess it would be semantically correct

Pro tip:
__when someone says "binary file", hit him in the face__
But i'm not sure. Maybe I just don't understand something here.









https://code.quora.com/Is-a-zip-file-a-binary-file