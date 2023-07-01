all:
	cd operations && python gen_op_codes.py
	python gen_access_flags.py
	scalac *.scala analysis/*.scala attributes/*.scala operations/*.scala test/*.scala

test-rewrite: clean all
	mkdir test_rewrite
	cp *.class test_rewrite
	scala RewriteClassFile test_rewrite/*.class

clean:
	rm -f *.class
	rm -f analysis/*.class
	rm -f attributes/*.class
	rm -f operations/*.class
	rm -f test/*.class
	rm -rf test_rewrite
