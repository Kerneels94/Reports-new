run:
	@cd view/auth && templ generate
	@cd ..
	@cd ..
	@cd cmd && air