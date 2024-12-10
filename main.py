f, s, p = [], [], 0
for i, l in enumerate( map( int, open( "input.txt" ).read().strip() ) ):
    ( f, s )[ i % 2 ].append( ( p, l ) )
    p += l

for fi in range( len( f ) - 1, -1, -1 ):
    fp, fl = f[ fi ]
    for si, ( sp, sl ) in enumerate( s ):
        if sl >= fl:
            f[ fi ] = ( sp, fl )
            s[ si ] = ( sp + fl, sl - fl )
            break
        if sp >= fp:
            break

print( sum( sum( n * x for x in range( p, p + l ) )
            for n, ( p, l ) in enumerate( f ) ) )