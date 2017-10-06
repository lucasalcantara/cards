$(document).ready(function(){
	
	$("#create").click(function(){

		var json = [{
			"question": $("#question").val(),
			"answer": $("#answer").val(),
		}];


		$.ajax({
			type: "POST",
			url: "new/create",
			data: JSON.stringify(json),
			success: function(data){
				alert(data);
			}
		});
	});


	$(".card-grid").flip({
		trigger: 'manual'
	});

	$(".card-grid button[class=speech]").click(function(){
		var text = $(this).parent().find('p[class=phrase]').text();
		var adjustedText = adjustText(text);

		var url = "https://translate.google.com/translate_tts?ie=UTF-8&q="+ encodeURI(adjustedText + "!") +"&tl=cs-CZ&client=tw-ob";

		// TODO: find a better way to do the listen.
	    var promise = $('audio').attr("src", url).get(0);
	    promise.play().catch(function(){
			url = "https://translate.google.com/translate_tts?ie=UTF-8&q="+ encodeURI(adjustedText + "!") +"&tl=cs&client=tw-ob";
			$('audio').attr("src", url).get(0).play().catch(function(){
				url = "https://translate.google.com/translate_tts?ie=UTF-8&q="+ encodeURI(adjustedText) +"&tl=cs&client=tw-ob";
				$('audio').attr("src", url).get(0).play().catch(function(){
					url = "https://translate.google.com/translate_tts?ie=UTF-8&q="+ encodeURI(adjustedText) +"&tl=cs-CZ&client=tw-ob";
					$('audio').attr("src", url).get(0).play();
				});
			});
	    });
	});

	function adjustText(text){
		
		text = text.replace("?", "")
		return text.replace(",", " ").toLowerCase();
	} 

	$(".card-grid button[class=close]").click(function(){

		$.ajax({
			type: "POST",
			url: "remove",
			data: "id=" + $(this).parent().find("input").val(),
			success: function(data){
				window.location.href = window.location.href;
			}
		});
	});

		
	$(".card-grid div").click(function(e){
		 if (e.target.tagName === "BUTTON" || e.target.tagName === "INPUT")
		    return;
		
		passAnswerToBackCard($(this));
		$(this).parent().flip('toggle');
	});

	function passAnswerToBackCard(e){
		if(e.parent().find("input[type=text]").length){
			var text = e.parent().find("input[type=text]").val();
			e.parent().find("input[type=text]").val("");
			e.parent().find("div[class=back]").find("p[class=compare]").text(text);
		}
	}
});