# shape Area

Here is an n-interesting polygon. We must find the area of a polygon given for n.

a 1 interesting polygon is just a square with a length of one.

1 = []

An n-interesting pologon is obtained by taking n - 1-interesting pologygon and appending 1-interesting polygon to the sides of each open side.

   [] n = 1

         []
       [][][] n = 2
         []

            []
          [][][]
        [][][][][] n = 3
          [][][]
            []

               []
             [][][]
           [][][][][]
         [][][][][][][] n = 4
           [][][][][]
             [][][]
               []

input n = 2
output = 5

input n = 3
output = 13

input n = 4
output = 25
