# circleci webhoook to ikachan

This is a forwarder from circleci webhook to ikachan

## SYNOPSIS

    go build

     ./circleci_ikachan -ikachan http://localhost:5000

## Sending test request

    curl -X POST -d @test.json http://localhost:3000\?channel\=%23nekokak\&message_type\=notice

## Ikachan dummy server for debugging

You may need ikachan dummy server for debugging...

You can use plackup for testing.

    plackup -e 'sub {use Plack::Request; my $env = shift; warn Plack::Request->new($env)->content; [200, [], []] }'

## LICENSE

The MIT License (MIT)
Copyright © 2016 Tokuhiro Matsuno, http://64p.org/ <tokuhirom@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

