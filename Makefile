run:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./
	zip myFunction.zip bootstrap  
	aws lambda update-function-code --function-name $m --zip-file fileb://myFunction.zip

deploy:
	make run m=dev-check-public-ip

commit:
	git add .
	git commit -m "$m"
	git push origin main