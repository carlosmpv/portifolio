$ document
	.ready ->
		handle_smooth($('a[href*="#"]').not('[href="#"]').not('[href="#0"]'))
		handle_comment($("#comment-form"))
		handle_navbar($("#navbarNav"))

		load_relevant_commentaries()
