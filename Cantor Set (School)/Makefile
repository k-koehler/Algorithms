tests:
	@go build -o test.run
	@./test.run < test0.in > test0.out
	@./test.run < test1.in > test1.out
	@echo 'Done tests.'
	@echo 'Difference between expected output and output for test0 (if nothing displays then no difference)'
	@diff test0.exp test0.out
	@echo 'Difference between expected output and output for test1 (if nothing displays then no difference)'
	@diff test1.exp test1.out
	
