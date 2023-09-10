---- Sep 09 16:11:22 +07 2023
# How I failed to understand base64

## Base64 is astonishingly difficult to understand. I mean, the algorithm is dead simple and it won't be discussed in this text. I rather tried to understand 1) in what situations this algorithm is useful and 2) how exactly it solves the situation. It's striking how many varying opinions exist about the purpose of base64. Official documents are very short and dry on this subject. Apparently it's all buried in history and digging it won't be extremely rewarding... But still it bugs me that there is not a single coherent explanation of such a widespread technology. I tried to create such an explanation but failed.


Convention
............
Typically all discussions start from (and often fail due to) a statement like:
    "Base64 is a way to transform binary data into text".
Such statements sound kinda reasonable. Because there is a convention that binary data !== text. But in my opinion this convention should be revised. (I wrote about it in my previous article).
Saying that "binary is transformed into text" is rather untrue and surely very unhelpful. It leads your brain in the wrong direction. Because you can't really transform binary into text. You can only transform binary into another binary. And that's what base64 does, it takes one sequence of bits and outputs another sequence of bits. If we adhere to this idea, it will be easier to understand further stuff.
Therefore I don't say "binary data", I say just "data".
............

### Why

Let's start with popular explanations.

There are many schools of thought regarding "Why we need base64" (In response to what it was introduced):
1) Because some system can deal with ascii characters only.
2) Because some systems think that byte consists of 7 bits, not 8.
3) Because some data can contain some dangerous characters that can be misinterpreted.
Apparently all these are real-world problems. But is base64 a solution to them? I'm reluctant to accept these explanations because most of them leave an impression of utter incompetence. So I'd like to check all of these theories.

### When

In what situations we may need base64?
When the aforementioned 3 problems may arise?
Wikipedia in 2023 says: when we TRANSMIT data "across channels that only reliably support text content".
Transmission is indeed the most believable case. But not the only one.
RFC4648 says that base 64 is used
"...to store or transfer data in environments that are restricted to ASCII data"
So we may need it when we STORE something too. But it's really difficult to say what they mean. Store what, where and when? RFC gives no further clarification. Let's pretend we didn't read it.
RFC continues that base64 is also used "...to make it possible to manipulate objects with text editors". This thread is also very difficult to develop.

### Where

We  know that base64 was introduced by SMTP (email protocol). But it surely wasn't the only user of base64. There must be a whole class of text-based systems that were in question. But I failed to find any information about this class.

In short, base64 is an obscure territory that disappears in history. Surely everything was very complicated, and base64 solved a whole bunch of problems in whole bunch of different systems. Digging this history seems a very unrewarding enterprise.


### Let's at least try to understand the case of SMTP

SMTP is big thing. (Still a primary protocol for sending email messages today. Used by Gmail, Outlook etc).
It has at least two of the aforementioned problems: It allows only ASCII characters and it reads data in 7-bit chunks.

### What this ASCII restriction means precisely?

All data transmitted via SMTP protocol must contain ONLY characters from ASCII set. And ASCII is just 128 characters. And it allows for communication only in English language.
If SMTP-compliant email client finds something that doesn't represent ASCII character,
various things may happen. Data can be lost (Rejected) or modified (e.g., an attempt to convert to ASCII can be made).

### How this "ASCII only" restriction works at a binary level?

At a byte level this restriction means that SMPT-compliant systems can transmit only such data where every 7 bits could be translated into an ASCII character.
So for instance a sequence 124124124124 could travel via SMTP because it translates to "...". And 52095839 couldn't because it's doesn't translate to any of ASCII characters.
So the problem with SMTP is double-sided:
1) it expects a very limited set of characters
2) it looks for characters in the "wrong" way. It assumes that every character is expressed by 7 bits of data. Whereas in modern systems it's more common to use 8-bit chunks.
Can base64 help with these problems? Maybe. Read further.

### Can base64 solve "ASCII only" problem?

ø你Ω
Apparently it can. It's one of the few things we can be (almost) sure of.

Actually base64 was only a PART of a bigger solution which is MIME.
Base64 is a very simple thing and it makes sense only as part of a more complex mechanism that uses this encoding for good.

Let's look at it in detail.
1) A byte sequence that contains non-ascii
2) Encode it to base64 and get a string of ASCII characters
3) Examine the binary content of this string
4) Make sure that this binary content doesn't contain "0000" any more.





Let's say we have the following binary sequence:
01000001 00000000 01000010.
First byte, when translated to ASCII, means "A";
then goes null character;
then ASCII character "B"
In short, this sequence means "A\0B".

Now look at the gibberish generated by base64:

```
printf "A\0B" | base64 -w 0
// QQBC
```

Now here's the binary representation of this gibberish:
```
echo -n -e "QQBC" | xxd -b
// 01010001 01010001 01000010 01000011
```

Looks good:
1) after base64 encoding we have a different binary sequence;
2) this new sequence doesn't contain null character.

We can rest assured that base64 encoded data is unambiguous. This data is ready to be safely transmitted / processed by any systems. (But I'm only guessing).

# But...
## How exactly do we use this base64-encoded stuff?

It seems that base64-encoded data is just gibberish.
This data is NOT an original data minus detrimental characters.
Base64 doesn't do anything very clever actually. It's unaware of dangerous characters. It's also unaware of what the original bytes stood for.
All that it does is take 6 bits of original data and replace it with ASCII character, then next 6 bits and the next 6...
In the end you get just a sequence of characters which are safe but basically useless. ...Until we decode it back to original form! I.e., an original form that contained "0000" and anything that was dangerous there in the beginning.

For example, when we encode an image to base64, it's no longer an image. I.e. this base64-encoded data cannot be USED as an image, it can't be displayed. To display it we need to convert base64 data back to the original bytes. Original bytes still contain ambiguous characters. But this time for some reason... they are no longer dangerous!





















### It's often said that base64 helps with some ambiguous binary sequences.


-------------------------------
That is, for instance, a sequence "00000000", or "a null character". It is said that such character can be treated as a signal that  "data ends here"
Another popular example is "1010", or a "Line feed" character from an ASCII point of view.
People say that such characters can be misinterpeted by some systems. And that base64 somehow helps with such characters.
There is some logic in it. Base64 uses only this limited alphabet of 64 symbols because they are absolutely safe and will treated in the same way by overwelming majority of systems.
Indeed when data is encoded to base64, "00000000" has no chance. It will be converted to some other binary sequence representing ASCII chars.

But what's the use of it?
In a "contraband" scenario it makes little sense.
step by step
1. Let's imagine a system A that has a data to send to system B. Data contains line feed characters that, from point of view of system A, indicate a new line. But system B interprets line feed character somewhat differently, suppose it ignores it.
2. let's encode it to base64.
3. Now we have a base64 gibberish. It's unusable as text. If the receiving system is going to display it as text, it needs to decode it back to original data.
4. Decoded data is a precise copy of what we had in the beginning. I.e., it contains same line feed character. So if a receiving system can't interpret this character correctly, it will not interpret correctly Encoding-decoding this data using base64 will make no difference.

**************
does SMTP actully read files while transmitting?
(and in the process of reading, can a null character stop reading and halt transmission?)












It's used for file uploads















to read
    "Why "optimizing" your images with Base64 is almost always a bad idea"











SMTP RFC:
"If the transmission channel provides an 8-bit byte (octet) data
   stream, the 7-bit ASCII codes are transmitted, right justified, in
   the octets, with the high-order bits cleared to zero. "

SMTP suggests a way to send 7bit data across 8bit channel. It's easy. If recipient says it's 8bit, you just add 0.
So you don't need different versions of base64. You have only 1 that emits 7bit data and then, if necessary, you adjust this data with 0s.