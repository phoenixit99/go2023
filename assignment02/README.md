# go2023

<h4>Assignment02</h4>

Sort func by input multiple type </br>

Create a package and a
command-line tool to sort input provided
by the user.</br>
Inputs: Number (integer or float) array,
string array.</br>
Outputs: Sorted result based on the
provided input type.


<pre>
Example #1 : 

Inputs: go main -int 2 1 5 4 1

Output: 1 1 2 4 5


 
  
Example #2 : 

Inputs: go run main.go -mix 2.3 21 123 1 31.1 12 golang apple

Output: 1 2.3 12 21 31.1 123 apple golang 

Example #3 : 

Inputs:  go run main.go -string text a juce go  

Output: a go juce text
</pre>
  
<h4>How to work: </h4>
 <pre>
--- Run the command line --- </br>
go run main.go [flag] (flag: -int, -string, -mix)
</pre>
