package server

const indexHTML = `<!DOCTYPE html>
<html>
<head>
    <meta charSet="utf-8" />
    <meta name="viewport" content="user-scalable=no,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0" />
    <title>ESM</title>
    <meta name="description" content="A fast, global content delivery network and package manager for ES Modules." />
    <meta name="keywords" content="esm,npm,deno,global,cdn" />
    <style>
        /*
         * Name: Github ReadMe style for Mou app
         * Version: v1.1
         * Author: hzlzh(hzlzh.dev@gmail.com)
         * URL: https://github.com/hzlzh/Mou-Theme
         */

        @charset "UTF-8";

        html {
            font-family: sans-serif;
            -ms-text-size-adjust: 100%%;
            -webkit-text-size-adjust: 100%%
        }

        body {
            margin: 0
        }

        article,
        aside,
        details,
        figcaption,
        figure,
        footer,
        header,
        hgroup,
        main,
        nav,
        section,
        summary {
            display: block
        }

        audio,
        canvas,
        progress,
        video {
            display: inline-block;
            vertical-align: baseline
        }

        audio:not([controls]) {
            display: none;
            height: 0
        }

        [hidden],
        template {
            display: none
        }

        a {
            background: transparent
        }

        a:active,
        a:hover {
            outline: 0
        }

        abbr[title] {
            border-bottom: 1px dotted
        }

        b,
        strong {
            font-weight: bold
        }

        dfn {
            font-style: italic
        }

        h1 {
            font-size: 2em;
            margin: 0.6em 0
        }

        mark {
            background: #ff0;
            color: #000
        }

        small {
            font-size: 80%%
        }

        sub,
        sup {
            font-size: 75%%;
            line-height: 0;
            position: relative;
            vertical-align: baseline
        }

        sup {
            top: -0.5em
        }

        sub {
            bottom: -0.25em
        }

        img {
            border: 0
        }

        svg:not(:root) {
            overflow: hidden
        }

        figure {
            margin: 1em 40px
        }

        hr {
            -moz-box-sizing: content-box;
            box-sizing: content-box;
            height: 0
        }

        pre {
            overflow: auto
        }

        code,
        kbd,
        pre,
        samp {
            font-family: 'source code pro', monospace;
            font-size: 1em
        }

        button,
        input,
        optgroup,
        select,
        textarea {
            color: inherit;
            font: inherit;
            margin: 0
        }

        button {
            overflow: visible
        }

        button,
        select {
            text-transform: none
        }

        button,
        html input[type="button"],
        input[type="reset"],
        input[type="submit"] {
            -webkit-appearance: button;
            cursor: pointer
        }

        button[disabled],
        html input[disabled] {
            cursor: default
        }

        button::-moz-focus-inner,
        input::-moz-focus-inner {
            border: 0;
            padding: 0
        }

        input {
            line-height: normal
        }

        input[type="checkbox"],
        input[type="radio"] {
            box-sizing: border-box;
            padding: 0
        }

        input[type="number"]::-webkit-inner-spin-button,
        input[type="number"]::-webkit-outer-spin-button {
            height: auto
        }

        input[type="search"] {
            -webkit-appearance: textfield;
            -moz-box-sizing: content-box;
            -webkit-box-sizing: content-box;
            box-sizing: content-box
        }

        input[type="search"]::-webkit-search-cancel-button,
        input[type="search"]::-webkit-search-decoration {
            -webkit-appearance: none
        }

        fieldset {
            border: 1px solid #c0c0c0;
            margin: 0 2px;
            padding: 0.35em 0.625em 0.75em
        }

        legend {
            border: 0;
            padding: 0
        }

        textarea {
            overflow: auto
        }

        optgroup {
            font-weight: bold
        }

        table {
            border-collapse: collapse;
            border-spacing: 0
        }

        td,
        th {
            padding: 0
        }

        * {
            -moz-box-sizing: border-box;
            box-sizing: border-box
        }

        input,
        select,
        textarea,
        button {
            font: 13px/1.4 Helvetica, arial, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol"
        }

        body {
            padding: 30px;
            max-width: 900px;
            margin: 0 auto;
            font: 14px/1.4 Helvetica, arial, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol";
            color: #333;
            background-color: #fff
        }

        a {
            color: #0366d6;
            text-decoration: none
        }

        a:hover,
        a:focus,
        a:active {
            text-decoration: underline
        }

        hr,
        .rule {
            height: 0;
            margin: 15px 0;
            overflow: hidden;
            background: transparent;
            border: 0;
            border-bottom: 1px solid #ddd
        }

        hr:before,
        .rule:before {
            display: table;
            content: ""
        }

        hr:after,
        .rule:after {
            display: table;
            clear: both;
            content: ""
        }

        fieldset {
            padding: 0;
            margin: 0;
            border: 0
        }

        label {
            font-size: 13px;
            font-weight: bold
        }

        input[type="text"],
        #adv_code_search .search-page-label,
        input[type="password"],
        input[type="email"],
        input[type="number"],
        input[type="tel"],
        input[type="url"],
        textarea {
            min-height: 34px;
            padding: 7px 8px;
            font-size: 13px;
            color: #333;
            vertical-align: middle;
            background-color: #fff;
            background-repeat: no-repeat;
            background-position: right center;
            border: 1px solid #ccc;
            border-radius: 5px;
            outline: none;
            box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.075)
        }

        input[type="text"].focus,
        #adv_code_search .focus.search-page-label,
        input[type="text"]:focus,
        .focused .drag-and-drop,
        #adv_code_search .search-page-label:focus,
        input[type="password"].focus,
        input[type="password"]:focus,
        input[type="email"].focus,
        input[type="email"]:focus,
        input[type="number"].focus,
        input[type="number"]:focus,
        input[type="tel"].focus,
        input[type="tel"]:focus,
        input[type="url"].focus,
        input[type="url"]:focus,
        textarea.focus,
        textarea:focus {
            border-color: #51a7e8;
            box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.075), 0 0 5px rgba(81, 167, 232, 0.5)
        }

        input.input-contrast,
        .input-contrast {
            background-color: #fafafa
        }

        input.input-contrast:focus,
        .input-contrast:focus {
            background-color: #fff
        }

        ::-webkit-input-placeholder,
        :-moz-placeholder {
            color: #aaa
        }

        ::-webkit-validation-bubble-message {
            font-size: 12px;
            color: #fff;
            background: #9c2400;
            border: 0;
            border-radius: 5px;
            -webkit-box-shadow: 1px 1px 1px rgba(0, 0, 0, 0.1)
        }

        input::-webkit-validation-bubble-icon {
            display: none
        }

        ::-webkit-validation-bubble-arrow {
            background-color: #9c2400;
            border: solid 1px #9c2400;
            -webkit-box-shadow: 1px 1px 1px rgba(0, 0, 0, 0.1)
        }

        body {
            font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif;
            font-size: 16px;
            line-height: 1.5;
            word-wrap: break-word
        }

        .absent {
            color: #c00
        }

        .anchor {
            position: absolute;
            top: 0;
            bottom: 0;
            left: 0;
            display: block;
            padding-right: 6px;
            padding-left: 30px;
            margin-left: -30px
        }

        .anchor:focus {
            outline: none
        }

        h1,
        h2,
        h3,
        h4,
        h5,
        h6 {
            position: relative;
            margin-top: 1em;
            margin-bottom: 16px;
            font-weight: bold;
            line-height: 1.4
        }

        h1 .octicon-link,
        h2 .octicon-link,
        h3 .octicon-link,
        h4 .octicon-link,
        h5 .octicon-link,
        h6 .octicon-link {
            display: none;
            color: #000;
            vertical-align: middle
        }

        h1:hover .anchor,
        h2:hover .anchor,
        h3:hover .anchor,
        h4:hover .anchor,
        h5:hover .anchor,
        h6:hover .anchor {
            height: 1em;
            padding-left: 8px;
            margin-left: -30px;
            line-height: 1;
            text-decoration: none
        }

        h1:hover .anchor .octicon-link,
        h2:hover .anchor .octicon-link,
        h3:hover .anchor .octicon-link,
        h4:hover .anchor .octicon-link,
        h5:hover .anchor .octicon-link,
        h6:hover .anchor .octicon-link {
            display: inline-block
        }

        h1 tt,
        h1 code,
        h2 tt,
        h2 code,
        h3 tt,
        h3 code,
        h4 tt,
        h4 code,
        h5 tt,
        h5 code,
        h6 tt,
        h6 code {
            font-size: inherit
        }

        h1 {
            padding-bottom: 0.3em;
            font-size: 2.25em;
            line-height: 1.2;
            border-bottom: 1px solid #eee
        }

        h2 {
            padding-bottom: 0.3em;
            font-size: 1.75em;
            line-height: 1.225;
            border-bottom: 1px solid #eee
        }

        h3 {
            font-size: 1.5em;
            line-height: 1.43
        }

        h4 {
            font-size: 1.25em
        }

        h5 {
            font-size: 1em
        }

        h6 {
            font-size: 1em;
            color: #777
        }

        p,
        blockquote,
        ul,
        ol,
        dl,
        table,
        pre {
            margin-top: 0;
            margin-bottom: 12px
        }

        hr {
            height: 4px;
            padding: 0;
            margin: 16px 0;
            background-color: #e7e7e7;
            border: 0 none
        }

        ul,
        ol {
            padding-left: 2em
        }

        ul.no-list,
        ol.no-list {
            padding: 0;
            list-style-type: none
        }

        ul ul,
        ul ol,
        ol ol,
        ol ul {
            margin-top: 0;
            margin-bottom: 0
        }

        li>p {
            margin-top: 16px
        }

        dl {
            padding: 0
        }

        dl dt {
            padding: 0;
            margin-top: 16px;
            font-size: 1em;
            font-style: italic;
            font-weight: bold
        }

        dl dd {
            padding: 0 16px;
            margin-bottom: 16px
        }

        blockquote {
            padding: 0 15px;
            color: #777;
            border-left: 4px solid #ddd
        }

        blockquote> :first-child {
            margin-top: 0
        }

        blockquote> :last-child {
            margin-bottom: 0
        }

        table {
            display: block;
            width: 100%%;
            overflow: auto;
            word-break: normal;
            word-break: keep-all
        }

        table th {
            font-weight: bold
        }

        table th,
        table td {
            padding: 6px 13px;
            border: 1px solid #ddd
        }

        table tr {
            background-color: #fff;
            border-top: 1px solid #ccc
        }

        table tr:nth-child(2n) {
            background-color: #f8f8f8
        }

        img {
            max-width: 100%%;
            -moz-box-sizing: border-box;
            box-sizing: border-box;
            border-radius: 5px;
        }

        span.frame {
            display: block;
            overflow: hidden
        }

        span.frame>span {
            display: block;
            float: left;
            width: auto;
            padding: 7px;
            margin: 13px 0 0;
            overflow: hidden;
            border: 1px solid #ddd
        }

        span.frame span img {
            display: block;
            float: left
        }

        span.frame span span {
            display: block;
            padding: 5px 0 0;
            clear: both;
            color: #333
        }

        span.align-center {
            display: block;
            overflow: hidden;
            clear: both
        }

        span.align-center>span {
            display: block;
            margin: 13px auto 0;
            overflow: hidden;
            text-align: center
        }

        span.align-center span img {
            margin: 0 auto;
            text-align: center
        }

        span.align-right {
            display: block;
            overflow: hidden;
            clear: both
        }

        span.align-right>span {
            display: block;
            margin: 13px 0 0;
            overflow: hidden;
            text-align: right
        }

        span.align-right span img {
            margin: 0;
            text-align: right
        }

        span.float-left {
            display: block;
            float: left;
            margin-right: 13px;
            overflow: hidden
        }

        span.float-left span {
            margin: 13px 0 0
        }

        span.float-right {
            display: block;
            float: right;
            margin-left: 13px;
            overflow: hidden
        }

        span.float-right>span {
            display: block;
            margin: 13px auto 0;
            overflow: hidden;
            text-align: right
        }

        code,
        tt {
            padding: 0;
            padding-top: 0.2em;
            padding-bottom: 0.2em;
            margin: 0;
            font-size: 85%%;
            background-color: rgba(0, 0, 0, 0.04);
            border-radius: 5px
        }

        code:before,
        code:after,
        tt:before,
        tt:after {
            letter-spacing: -0.2em;
            content: "\00a0"
        }

        code br,
        tt br {
            display: none
        }

        del code {
            text-decoration: inherit;
            vertical-align: text-top
        }

        pre>code {
            padding: 0;
            margin: 0;
            font-size: 100%%;
            word-break: normal;
            white-space: pre;
            background: transparent;
            border: 0
        }

        .highlight {
            margin-bottom: 16px
        }

        .highlight pre,
        pre {
            padding: 16px;
            overflow: auto;
            font-size: 85%%;
            line-height: 1.45;
            background-color: #f8f8f8;
            border-radius: 5px
        }

        .highlight pre {
            margin-bottom: 0;
            word-break: normal
        }

        pre {
            word-wrap: normal
        }

        pre code,
        pre tt {
            display: inline;
            max-width: initial;
            padding: 0;
            margin: 0;
            overflow: initial;
            line-height: inherit;
            word-wrap: normal;
            background-color: transparent;
            border: 0
        }

        pre code:before,
        pre code:after,
        pre tt:before,
        pre tt:after {
            content: normal
        }
    </style>
    <style>
        /* esm.sh */

        h1, h2, h3 {
            color: #111;
            font-weight: 500;
        }

        h1 {
            margin-top: 2.1em;
            font-size: 2.1em;
            font-weight:600;
        }

        h2 {
            font-size: 1.5em;
        }

        h3 {
            font-size: 1.2em;
        }

        h1 strong {
            display: inline-block;
            vertical-align: middle;
        }

        h1 a {
            position: relative;
            display: inline-block;
            vertical-align: middle;
            width: 27px;
            height: 27px;
            margin-left: 12px;
            color: #555;
        }

        h1 a svg {
            position: absolute;
            display: inline-block;
            top: 0;
            left: 0;
            width: 27px;
            height: 27px;
        }

        h1 a:hover {
           color: #000;
        }

        .test {
            display: none;
        }

        .test p {
            display: inline-block;
        }
    </style>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@10.1.2/build/styles/github.min.css" />
</head>
<body>
    <h1>
        <strong>ESM</strong>
        <a href="https://github.com/postui/esm.sh">
            <svg fill="currentColor" viewBox="0 0 24 24">
                <title>Github</title>
                <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd"></path>
            </svg>
        </a>
    </h1>
    <main>
        <p><em style="color: #999;">Loading...</em></p>
    </main>
    <div class="test">
        <h2>Test</h2>
        <p><strong>React</strong>: <span id="react">❌</span></p>
        <p><strong>React(bundle)</strong>: <span id="reactb">❌</span></p>
        <p><strong>Preact</strong>: <span id="preact"><span className="x">❌</span></span></p>
        <p><strong>Vue2</strong>: <span id="vue">❌</span></p>
        <p><strong>Vue3</strong>: <span id="vue3">❌</span></p>
    </div>
    <script type="module">
        import marked from '/marked';
        import hljs from '/highlight.js/lib/core';
        import javascript from '/highlight.js/lib/languages/javascript';
        import json from '/highlight.js/lib/languages/json';
        import bash from '/highlight.js/lib/languages/bash';

        const mainEl = document.querySelector('main');
        mainEl.innerHTML = marked.parse(%s);
        mainEl.removeChild(mainEl.querySelector('h1'));
        hljs.registerLanguage('javascript', javascript);
        hljs.registerLanguage('json', json);
        hljs.registerLanguage('bash', bash);
        hljs.initHighlighting();

        mainEl.querySelectorAll('img').forEach(img => {
            const src = img.getAttribute('src')
            if (src.startsWith('./assets/')) {
                img.src = 'https://raw.githubusercontent.com/postui/esm.sh/master/' + src.replace('./', '')
            }
        })

        if (location.hostname === 'localhost') {
            document.querySelector('.test').style.display = "block";
        }
    </script>
    <script type="module">
        import { createElement } from '/react';
        import { render } from '/react-dom';

        render(createElement('span', null, '✅'), document.querySelector('#react'));
    </script>
    <script type="module">
        import { createElement } from '/[react,react-dom]/react';
        import { render } from '/[react,react-dom]/react-dom';

        render(createElement('span', null, '✅'), document.querySelector('#reactb'));
    </script>
    <script type="module">
        import { h, render } from '/preact';
        import { useEffect } from '/preact/hooks';

        const el = document.querySelector('#preact');
        function App() {
            useEffect(() => {
                el.removeChild(el.querySelector('.x'))
            }, [])
            return h('span', null, '✅')
        }
        render(h('span', null, '✅'), el);
    </script>
    <script type="module">
        import Vue from '/vue';

        new Vue({
            el: '#vue',
            render(h) {
                return h('span', null, '✅')
            }
        });
    </script>
    <script type="module">
        import { createApp, h } from '/vue@3.0.0';

        createApp({
            render() {
                return h('span', {}, '✅')
            }
        }).mount('#vue3');
    </script>
</body>
</html>
`

// copy from https://github.com/defunctzombie/node-process/blob/master/browser.js
const processBrowserJS = `// shim for using process in browser
var process = {};

// cached from whatever global is present so that test runners that stub it
// don't break things.  But we need to wrap it in a try catch in case it is
// wrapped in strict mode code which doesn't define any globals.  It's inside a
// function because try/catches deoptimize in certain engines.

var cachedSetTimeout;
var cachedClearTimeout;

function defaultSetTimeout() {
    throw new Error('setTimeout has not been defined');
}

function defaultClearTimeout () {
    throw new Error('clearTimeout has not been defined');
}

(function () {
    try {
        if (typeof setTimeout === 'function') {
            cachedSetTimeout = setTimeout;
        } else {
            cachedSetTimeout = defaultSetTimeout;
        }
    } catch (e) {
        cachedSetTimeout = defaultSetTimeout;
    }
    try {
        if (typeof clearTimeout === 'function') {
            cachedClearTimeout = clearTimeout;
        } else {
            cachedClearTimeout = defaultClearTimeout;
        }
    } catch (e) {
        cachedClearTimeout = defaultClearTimeout;
    }
} ())

function runTimeout(fun) {
    if (cachedSetTimeout === setTimeout) {
        //normal enviroments in sane situations
        return setTimeout(fun, 0);
    }
    // if setTimeout wasn't available but was latter defined
    if ((cachedSetTimeout === defaultSetTimeout || !cachedSetTimeout) && setTimeout) {
        cachedSetTimeout = setTimeout;
        return setTimeout(fun, 0);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedSetTimeout(fun, 0);
    } catch(e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't trust the global object when called normally
            return cachedSetTimeout.call(null, fun, 0);
        } catch(e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error
            return cachedSetTimeout.call(this, fun, 0);
        }
    }
}

function runClearTimeout(marker) {
    if (cachedClearTimeout === clearTimeout) {
        //normal enviroments in sane situations
        return clearTimeout(marker);
    }
    // if clearTimeout wasn't available but was latter defined
    if ((cachedClearTimeout === defaultClearTimeout || !cachedClearTimeout) && clearTimeout) {
        cachedClearTimeout = clearTimeout;
        return clearTimeout(marker);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedClearTimeout(marker);
    } catch (e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't  trust the global object when called normally
            return cachedClearTimeout.call(null, marker);
        } catch (e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error.
            // Some versions of I.E. have different rules for clearTimeout vs setTimeout
            return cachedClearTimeout.call(this, marker);
        }
    }
}

var queue = [];
var draining = false;
var currentQueue;
var queueIndex = -1;

function cleanUpNextTick() {
    if (!draining || !currentQueue) {
        return;
    }
    draining = false;
    if (currentQueue.length) {
        queue = currentQueue.concat(queue);
    } else {
        queueIndex = -1;
    }
    if (queue.length) {
        drainQueue();
    }
}

function drainQueue() {
    if (draining) {
        return;
    }
    var timeout = runTimeout(cleanUpNextTick);
    draining = true;

    var len = queue.length;
    while(len) {
        currentQueue = queue;
        queue = [];
        while (++queueIndex < len) {
            if (currentQueue) {
                currentQueue[queueIndex].run();
            }
        }
        queueIndex = -1;
        len = queue.length;
    }
    currentQueue = null;
    draining = false;
    runClearTimeout(timeout);
}

process.nextTick = function (fun) {
    var args = new Array(arguments.length - 1);
    if (arguments.length > 1) {
        for (var i = 1; i < arguments.length; i++) {
            args[i - 1] = arguments[i];
        }
    }
    queue.push(new Item(fun, args));
    if (queue.length === 1 && !draining) {
        runTimeout(drainQueue);
    }
};

// v8 likes predictible objects
function Item(fun, array) {
    this.fun = fun;
    this.array = array;
}
Item.prototype.run = function () {
    this.fun.apply(null, this.array);
};
process.title = 'browser';
process.browser = true;
process.env = {};
process.argv = [];
process.version = ''; // empty string to avoid regexp issues
process.versions = {};

function noop() {}

process.on = noop;
process.addListener = noop;
process.once = noop;
process.off = noop;
process.removeListener = noop;
process.removeAllListeners = noop;
process.emit = noop;
process.prependListener = noop;
process.prependOnceListener = noop;

process.listeners = function (name) { return [] }

process.binding = function (name) {
    throw new Error('process.binding is not supported');
};

process.cwd = function () { return '/' };
process.chdir = function (dir) {
    throw new Error('process.chdir is not supported');
};
process.umask = function() { return 0; };

process.env.NODE_ENV = "%s";
export default process;
`

var readmeMD string
