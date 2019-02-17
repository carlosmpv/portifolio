handle_comment = (form) ->
	form.submit (event) ->
		event.preventDefault()
		arrayData = $(this).serializeArray()
		jsonData = {}

		$.each arrayData, ->
			jsonData[this.name] = this.value || ''

		$.ajax {
			url: "/commentary/comment",
			method: "POST",
			data: JSON.stringify jsonData
			dataType: "json",
			success: ((data) ->
				load_relevant_commentaries()
				$(":input").val("")

				# console.log data
				return
			),
			error: ((xhr, resp, text) ->
				# console.log xhr, resp, text
				return
			)
		}

format_date = (date) ->
	important_part = date.slice(0, 10)
	year = important_part.slice(0, 4)
	month = important_part.slice(5, 7)
	day = important_part.slice(8, 10)

	# console.log month

	month_list = [
		"Jan.", "Fev.", "Mar.", "Abr.",
		"Mai.", "Jun.", "Jul.", "Ago.",
		"Set.", "Out.", "Nov.", "Dez."
	]

	return "#{day} de #{month_list[Number(month)]} de #{year}"

load_relevant_commentaries = ->
	$.ajax {
		url: "/commentary/top",
		method: "GET",
		success: ((data) ->
			$ "#top-commentaries"
				.text ""
			for item in data
				$ "#top-commentaries"
					.prepend """
<div class="row p-3 pr-5 m-3 text-dark commentary-box">
	<p class="bg-light pr-4 pl-2 pt-2 pb-3 w-100">
		#{item.Content}
	</p>
	<span class="px-1 text-faded">#{format_date(item.PostDate)}</span>
	<span class="ml-auto font-italic font-weight-bold bg-light author-name px-1">#{item.Author}</span>
</div>
					"""
				# console.log item
		)
	}
