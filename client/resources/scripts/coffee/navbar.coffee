handle_navbar = (navbar) ->
	section_names = navbar.find("a.nav-link")

	navbar_links = {}
	section_list = {}

	window_height = $(window).height()
	console.log window_height

	$.each section_names, (k, v) ->
		id = $(v).attr("href")
		navbar_links[id] = v
		section_list[id] = $(id)

	$(document).scroll ->
		page_scroll = $(this).scrollTop()

		$.each section_names, (k, v) ->
			id = $(v).attr("href")
			section_position = section_list[id].position().top
			section_height = section_list[id].height()

			if page_scroll + window_height/2 >= section_position && page_scroll < section_position + section_height/2
				$(navbar_links[id]).addClass("active")
			else
				$(navbar_links[id]).removeClass("active")
