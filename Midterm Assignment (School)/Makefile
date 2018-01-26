tests:
	@go build -o test.run
	@echo 'running test0'
	@./test.run < test_inputs/test0.in > test_inputs/test0.out
	@echo 'running test1'
	@./test.run < test_inputs/test1.in > test_inputs/test1.out
	@echo 'running test2'
	@./test.run < test_inputs/test2.in > test_inputs/test2.out
	@echo 'running test3'
	@./test.run < test_inputs/test3.in > test_inputs/test3.out
	@echo 'running test4'
	@./test.run < test_inputs/test4.in > test_inputs/test4.out
	@echo 'dif between test0.exp and test0.out (should produce no output)'
	@diff test_inputs/test0.exp test_inputs/test0.out
	@echo 'dif between test1.exp and test1.out (should produce no output)'
	@diff test_inputs/test1.exp test_inputs/test1.out
	@echo 'dif between test2.exp and test2.out (should produce no output)'
	@diff test_inputs/test2.exp test_inputs/test2.out
	@echo 'dif between test3.exp and test3.out (should produce no output)'
	@diff test_inputs/test3.exp test_inputs/test3.out
	@echo 'dif between test4.exp and test4.out (should produce no output)'
	@diff test_inputs/test4.exp test_inputs/test4.out
	
