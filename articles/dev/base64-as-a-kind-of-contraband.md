---- Sep 09 16:11:22 +07 2023
# Base64 as a kind of contraband

## Today base64 encoding is widely used and taken for granted. It's a way of transforming any data into a gibberish of digits and letters. We need it not for compression or security but... because in some places only such text-like gibberish is accepted. It seems that we have no choice but to use base64. At the same time it's a costly instrument that makes things a little slower. So I think we shouldn't use it blindly. Let's try to understand what problems it solves (or solved) and how. And why this encoding and not the other?


*The encoding algorithm of base64 is straightforward and not discussed here. You can find good explanations of it elsewhere in the internet*

---

### Why base64?

Base64 is a tricky subject to explain because it solves a whole range of problems in a whole range of systems. And many of these problems date back to 1980s. And finally, there are many variants of base64.

There are 3 main theories about the purpose of base64:

1) Because some systems are restricted to ASCII characters but are actually used for all kinds of data. Base64 can "camouflage" this data as ASCII and thus help this data to pass validation.

2) Because some older systems think that data consists of 7-bit chunks (bytes), whereas modern ones use 8-bit bytes. This may lead to misunderstanding between systems. And base64 presumably can help with this too - but it's not so obvious how.

3) Because some characters may have special meaning and this will differ from from system to system. Base64 tackles this by using only the most "purely textual" characters from ASCII set.

---

### When base64?

In what contexts can we face 3 problems mentioned above? Wikipedia in 2023 says that base64 is used "to __carry data across channels__ that only reliably support text content".

Alright, base64 may be necessary to __transmit__ data from one place to another. When data is transmitted, it has to satisfy the rules of a transmission protocol. Sometimes these rules are too strict and in such case base64 is used to bypass them.

But transmission is perhaps not the only case. RFC4648 (the official spec) says that base64 is used "...to __store__ or transfer data in environments that are restricted to ASCII data". But store what, where and when? RFC gives no further clarification.

It  continues that base64 "...makes it possible to manipulate objects with text editors". What text editors? This thread is also difficult to develop.

Hereafter base64 will be discussed mostly as an instrument of data transmission because all other contexts are a bit too obscure.

---

### A case of SMTP

Base64 was first introduced as a part of MIME which is a standard of sending email messages.

MIME at large solved various problems of SMTP, an age-old email protocol still prevalent today and used by Gmail, Outlook etc.

The original design of SMTP quickly turned out to be inconvenient. First of all, SMTP allowed only plain text in English language (such that consists of only 128 characters known as ASCII). Therefore some walkarounds were necessary to send 1) non-English text and 2) non-textual attachments.

MIME offered a standardized way to bypass the "ASCII only" restriction. The solution was to *encode* non-textual data as textual. Base64 was *one of* such encodings. Its algorithm splits the original data into chunks of 6 bits (without worrying about the meaning of these bits) and converts every chunk into a safe textual character. As a result any data begins to __look like__ text.

This is basically a very ugly practice that is somewhat similar to contraband. Data is obfuscated (well, really) in order to cheat a system that doesn't allow such kind of data.

And by the way restrictions that we bypass are meant for *safety* in the first place. But it was long ago and now we don't need this safety anymore. Instead we need an unobtrusive channel to transport anything whatsoever. So we are fighting against old rules which act like pesky bureaucracy. But that's what legacy systems are and we have to live with them.

---

### Are these limitations still relevant today?

Yes, despite all the extensions and tricks.

Many restrictions of SMTP are relaxed thanks to various extensions. For instance, "8BIT MIME" extension allows to send email messages in 8-bit bytes and in characters other than ASCII. (So in some cases base64 may not be necessary at all to send a letter with attachments).

But it's still impossible to ignore the old restrictions. Because there are outdated email servers which didn't adopt new extensions. And you have to be able to communicate with them even if your own email server supports all the modern stuff.

Before sending a message to a certain email server, you first ask what kind of SMTP rules it supports. E.g., does it implement the 8BIT MIME extension. If it doesn't, you probably need to convert your message to older format.

---

### How base64 helps with these limitations?

It's self-evident that base64 must be a solution to "ASCII only" problem because it transforms everything to ASCII characters. But it becomes less obvious when you combine it with "__7 bits__" problem. Because no matter what kind of characters you use, they must be somehow transmittable by both 7-bit and 8-bit channels, depending on the situation.

Experts usually say something like:
> "Base64 transforms binary sequence 01001101 01100001 01101110 *(whatever it means)* into text "TWFu".

Such statements can lead you to think that something binary becomes non-binary. (Because we have this strange convention that textual data is something opposite to binary data). In fact all ASCII characters produced by base64 are ultimately bits and bytes, just like the original data.

Here is a bash command to get a binary representation of a string:

`echo -n "TWFu" | xxd -b`

It will tell you that "TWFu" is actually "`01010100 01010111 01000110 01110101`". But if every character is 8 bits long, it won't pass validation by a 7-bit channel, and perhaps it won't be recognized as text at all. Apparently some additional binary juggling must take place to make it work for all channels.

Fortunately with ASCII characters this binary juggling is easy. Because to store an ASCII character in memory we actually need only 7 bits.
They can be fattened to 8 bits only if we need a conventional "octet". This is done simply by adding a "0" bit in front of the original 7 bits. For example, "T" can be stored as both `1010100` and `01010100`.

Therefore conversion between 7-bit and 8-bit ASCII characters is a matter of adding/removing the leftmost "0" bit. Apparently email servers have to perform this kind of stuff when talking to each other.

So let's keep in mind that base64 in itself doesn't solve the "7 bit" problem. It just produces ASCII characters and this allows for a brisk conversion between 7 and 8 bits. But this conversion is a responsibility of a wider system, such as MIME.


#### *Memory cost depends on byte length*

*By the way, if you use only 7 bits per character, then base64 must be less wasteful in terms of memory usage.*

*The main theory is that base64 causes a memory overhead of 33% (or 37% or whatever). But it seems to be correct only for 8-bit scenario.*

*Because, as previously explained, base64 converts every 6-bit chunk of original data into a single ASCII character. If such character is 8 bits long, it means that we are wasting 2 bits per every original chunk and it's about 33%, just as promised by the experts. But with 7-bit characters this loss must be about twice smaller.*

---

### Why only 64 characters?

Base64 would be an overkill if it was made only to bypass the ASCII restriction. 

ASCII is 128 characters long and if we used all of them, the encoding algorithm would be more memory-efficient. It's cumbersome to explain but, in short, with 128-long alphabet a single character could represent 7 (instead of only 6) bits of original data.

But the authors of base64 decided to use only *half* of ASCII characters, and by doing so they made base64 still more wasteful. Surely they had a very good reason for that.

In fact characters can be wrong (unsafe) for at least two reasons:
1) they can be __invalid__ (forbidden by a system, like SMTP forbids anything beyond ASCII)
2) they can be valid *but* mistakenly recognized as a __special character__

In ASCII there are 30+ "control characters". They are __not printable__ and are meant to cause __some other__ effects. For example: "line feed", "backspace", "delete", "escape". Many of these commands are a legacy of some prehistoric devices like teleprinters.

All non-printable characters certainly have to be excluded from encoding alphabet. So we are left with some 90+ printable ones. But printable doesn't mean safe and reliable. They can also have __special meaning__ in different systems. And a bunch of them have special meaning in SMTP, for example "@", "<", ">".

So it ended up with 64 chars - first of all, because 64 is easier to deal with algorithmically. And it looks like a really safe alphabet that can be used in a wide range of systems, not only SMTP. 

Unfortunately there are __only 62__ characters that are guaranteed to have no special meaning in all systems. They are digits, English small letters and English capital letters.
2 remaining characters are difficult to choose. The most popular candidates are "+" and "/" but still in some situations they will break stuff. For example, they have special meaning in URLs. That's why we have a "base64url" variant where the last two characters are "_" and "-".












#### *Base64 does not fix discrepancies between systems in how they interpret special characters*

*Some special characters are known to be highly __ambiguous__, that is interpreted differently by different systems. The most notorious are characters concerned with line breaks - "line feed" and "carriage return" (again, this is a heritage of prehistoric devices). Different systems have different opinions on how to combine these two chars to produce an actual line break.*

*There is a widespread __misconception__ that base64 somehow helps to reconcile these differences.*

*But base64 has nothing to do with how data is __interpreted__ - for example, how text is displayed on a screen. Because in order to display something you first need to __decode__ the base64 gibberish. Decoded data is obviously an exact copy of the original data. And it contains all the ambiguous characters that have been there before encoding. So, again, base64 can only __conceal__ special chars for the time of transmission; it doesn't magically sanitize data from dangerous chars.*

---


### HTTP/1.1 (or just "HTTP" for brevity) is a text-based protocol. Do we have to base64-encode all data to send it over HTTP?

A __body__ of HTTP message can be anything - it's not restricted to textual characters. So in most cases you don't have to encode anything, and all data can be sent in its original form, without jumbling its bits and bytes.

What is really "text-based" about HTTP is __headers__. Basically they are restricted to ASCII (it's not exactly true but it's a good practice to use only ASCII).

Today HTTP headers are used in many different ways - and sometimes they have to carry non-ASCII stuff too. For example, basic HTTP authentication scheme suggests that username and password should be sent in the "Authentication" header. Username and password can contain a lot of incompatible stuff and therefore have to be encoded to safe textual chars. Base64 is recommended to use in such cases. For this reason some developers think that base64 is a kind of encryption (data protection) which is not true.

---

### Why is base64 used for Data URLs? Can we use a more efficient encoding for this?

Data URLs are probably the most well-know use for base64 today. It's a way to inline various assets like images (not links to them but their actual code) into HTML or CSS files.

Note that Data URLs have nothing to do with *transmission* of data. We use base64 here __not__ because HTTP protocol forbids any binary sequences. When you send HTML or CSS file over network, HTTP doesn't care what's inside these files. But HTML and CSS files have to be only text to be properly __interpreted__ (displayed) by text editors and browsers.

It makes sense but still it's regrettable - because, again, base64 is expensive. This notorious 33% or 37% of memory overhead is especially annoying with Data URLs. In most cases it defeats their purpose entirely. The purpose is to improve performance of course. You get an image without extra HTTP request and thus save some milliseconds. But this performance gain is small and easily nullified by the extra bytes created by base64.

So why base64 is used for data URLs at all? Could we use some less wasteful encoding for that? (I.e. encoding that uses a wider alphabet and thus outputs shorter strings of characters).

At present it's impossible because browsers allow only __url-safe characters__ in Data URLs. And there aren't too much url-safe characters - just a tiny bit more than 64. Why are we restricted to url-safe characters? Because we insert the encoded data into places where browsers expect a URL.

In theory browsers could be smarter and relax this limitation when necessary. In such hypothetical case, is there a __better encoding__ for Data URLs? Or, to pose a broader question, - 

### Is there a better encoding for non-textual data inlined in HTML/CSS files?

At first glance, there must be a lot of space for improvement. Because HTML and CSS files can contain __any Unicode__ characters and it's more than 1 million. So it must be possible to find 256 characters to encode images, isn't it? Such encoding would have no (or almost no) memory overhead.

In practice everything is much more complicated. Why not encode stuff with chinese characters, for example? Because they are too heavy. They take 3 or even 4 bytes of memory. That's how UTF-8 encoding works - it uses different number of bytes for different characters. And we are interested only in 1-byte characters.

(We could consider using multi-byte characters for our baseNNN encoding but let's not go into that. We need only 1-byte characters, let's take it as an axiom).

How many 1-byte characters are there in UTF-8? Only 128, and it's a good old ASCII range. Can we take all of them? No, because, again, we need only printable characters. We really need to see them in text editors and dev tools and elsewhere. Then, just like in case of SMTP, we have to exclude a bunch of visible characters because they are special characters in HTML or CSS. For example, double quote (") won't do because it can prematurely terminate the Data URL string, like this: `
<img src="data:image/png;base64,iVBOR"w0KAA" />`

So a possible alphabet again shrinks to 80-90 characters. This in theory allows to create another encoding that will use slightly less memory than base64. Such encodings actually exist, for example base85 made in Adobe. It is more memory-efficient because it encodes 4 bytes of original data into 5 characters. But base85 is also much slower to compute so its overall benefits are tiny, if any. And by the way it's not intended for web development and contains characters that can break things in HTML and CSS. (Though it must be possible to build a similar but web-friendly algorithm by swapping some characters).

---

### Recap

Base64 was first introduced as a way to bypass a number of archaic restrictions imposed on email messages by SMTP protocol. Base64 allowed to camouflage any data as text in order to pass validation when transmitted between email servers. It also ensured that this pseudo-text contains only safe characters, i.e. 1) only printable ones and 2) only those that have no special meaning in SMTP and (hopefully) in most other systems.

The final alphabet happened to be really narrow and it allowed to use base64 practically everywhere (but in slightly different variants). It helps in great many cases where you have to mix textual and non-textual data. Modern systems also use base64 despite its significant memory cost. At first glance this practice looks strange because modern systems don't have these age-old restrictions as SMTP had. But it turns out that in most cases we still have very few *cheap* and *non-special* characters, and potential alternatives to base64 offer very small benefits.
