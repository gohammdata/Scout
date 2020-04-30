Current branch working in dev for parser feature.
<h1>Scout - The Tree Walking Interpreter for Schnauzer</h1>
<p>This Interpretor, Scout, implements all of the features of the programming language Schnauzer. Scout tokenizes and parses Schnauzer source code in a REPL, building up an internal representation of the code called by an abstract syntax tree and then evaluate the tree.</p>
<b>Scout's Major Parts</b>
<ul>
    <li>The Lexer</li>
    <li>The Parser</li>
    <li>The Abstract Syntax Tree (referred from now on by AST)</li>
    <li>The Internal Object System</li>
    <li>The Evaluator</li>
</ul>

<b>Scout is Programmed in Go WITHOUT dependencies or ANY 3rd Party Libraries</b>

<i>Optional Rant on "Why Go" and how great it is for your Reading Pleasure</i>
<p>Decision for bare Go - Go was created by (among others) Ken Thompson....you know that Turing award winning UNIX creator who with his partner changed the landscape of computers. That partner, Dennis Ritchie, created C which also changed everything (and was influenced by B written by Ken Thompson but I digress). My decision to write this interpreter (and the soon to come compiler which will run in my self-made VM) in Go is that it is a new language created and managed by Google with Ken Thompson as one of the many brains. IMHO Go is the next big programming language that will take over the Computer Science world (yes I see you RUST fan-boys out there scauffing). From C we received great programming languages and Operating Systems, and while Go is one of them; the foundation, cleanliness, and lack of need for dependencies and 3rd party frameworks in Go is why I believe we will see the same great outcome from Go itself.</p>
