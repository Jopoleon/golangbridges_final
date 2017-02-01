$(window).on('load', function () {
	$(".loader").fadeOut("slow");
});

$(document).ready(function () {
	$('#header1').fadeTo(0, 0);
	$('#header2').fadeTo(0, 0);
	$('#search1').fadeTo(0, 0);
	$('#search2').fadeTo(0, 0);
	$('#button12').attr('disabled', true);
	$('.hint1').fadeIn();
	$('.hint2').fadeTo(0, 0);


	$('input.chb').on('change', function () {
		$('input.chb').not(this).prop('checked', false);
	});

	$('input.chb').click(function () {
		if ($('input.chb').is(":checked")) {
			window.checkedValue = $(this).val();
			console.log(window.checkedValue);
			$('#button12').attr('disabled', false)
			$('.hint1').fadeTo('medium', 0);
			$('.hint2').animate({
				opacity: 1
			});

		} else {
			$('#button12').attr('disabled', true);
		}
	})
});

$(document).ready(function () {
	// Note, that I'm adding the argument event to .click function. It made for Mozila Browser.
	$("#button12").click(function (event) {
		console.clear();
		$('.hint2').fadeTo('medium', 0);
		event.preventDefault();
		console.log(window.checkedValue);
		//http://spun.fkpkzs.ru/Level/Gorny
		$.ajax({
			data: {
				//name: "some_name",
				"value": window.checkedValue
			},
			dataType: "json",
			type: "POST",
			//Node.js website API for getting waterlevel.
			//url: "https://floating-shore-25832.herokuapp.com/scrape",
			//localhost url for getting waterlevel from Go server
			//url: "http://localhost:3000/scrape",
			url: "https://golangbridges.herokuapp.com/scrape",


			//if post request is successful -> do function with logic
			success: function (data) {

				$('#header1').fadeTo('fast', 0);
				$('#header2').fadeTo('fast', 0);

				// checking the content of server response
				console.log(data)
				console.log("success " + data.time + " " + data.waterlevel + " " + data.shipHight);

				// Appending the content on HTML in div with id="resposnse21" 
				//                $("#response21")
				//                    .empty()
				//                    .append("Time: " + data.time + "<br>").hide().fadeIn()
				//                    .append("Waterlevel: " + data.waterlevel + "<br>").hide().fadeIn()
				//                    .append("Ship type: " + data.shipValue.value + "<br>").hide().fadeIn();
				window.waterlevel_time = data.time;
				window.waterlevel = data.waterlevel;
				window.shipValue = data.shipHight;

				//logic of sorting bridges array 

				//sample json array of bridges
				var obj = [
            
            {
            name: "Длинноесловонарусском123",
						height: 257
            },
            {
            name: "longwordinenglishhsdgad223",
						height: 343
            },
					{
						name: "Example in english",
						height: 230
                        },
					{
						name: "Пример названия на руссском",
						height: 270
                        },
					{
						name: "Кроткое",
						height: 400
                        },
					{
						name: "b1",
						height: 210
                        },
					{
						name: "b2",
						height: 300
                        },
             ]

				//initilizing new array for suitable bridges
				var obj_g = [];

				console.log(window.shipValue);
				console.log(window.waterlevel);

				//parsing string data respond from server to get integers
				var shipandwater = parseInt(window.shipValue) + parseInt(window.waterlevel);
				console.log(shipandwater);

				//sorting
				for (var i = 0; i < obj.length; i++) {
					
					if (obj[i].height > shipandwater)

					//new array of bridges that ship can pass under
					{
						
						obj_g.push(obj[i]);
					}

				}
				
				$("#table1").empty();

				$("#table2").empty();
				var fadeInduration = 300;

				$('#header1').animate({
					duration: 300,
					opacity: 1
				});

				var tbl1 = $("<table><tr class='t1'><th>Название</th><th>Высота(см)</th></tr>").attr("id", "all_bridges");
				//$("#table1").append(tbl1);
				$(tbl1).hide().appendTo("#table1").fadeIn(300);
				for (var i = 0; i < obj.length; i++) {
					var tr = "<tr class='t1'>";
					var td1 = "<td class='text-left1'>" + obj[i].name + "</td>";
					var td2 = "<td class='text-left1'>" + obj[i].height + "</td></tr>";

					$(tr + td1 + td2).hide().appendTo("#all_bridges").fadeIn(300);
					//$("#all_bridges").append(tr + td1 + td2).fadeIn();
				};

				$('#header2').animate({
					opacity: 1,
					duration: 300
				});
				//<th>Bridge Name</th><th>Bridge Height</th>
				var tbl2 = $("<table><tr class='t2'><th>Название</th><th>Высота(см)</th></tr>").attr("id", "vaildBridges");
				$(tbl2).hide().appendTo("#table2").fadeIn(300);
				for (var i = 0; i < obj_g.length; i++) {
					var tr = "<tr class='t2'>";
					var td1 = "<td class='text-left1'>" + obj_g[i].name + "</td>";
					var td2 = "<td class='text-left1'>" + obj_g[i].height + "</td></tr>";

					$(tr + td1 + td2).hide().appendTo("#vaildBridges").fadeIn(300);
				};

				//search in tables
				
				
				//second search
          
          $('#search1').animate({
					duration: 300,
					opacity: 1
				});
          $('#search2').animate({
					duration: 300,
					opacity: 1
				});
				
				$("#search1").keyup(function(){
        _this = this;
        // Show only matching TR, hide rest of them
        $.each($("#all_bridges .t1"), function() {
            if($(this).text().toLowerCase().indexOf($(_this).val().toLowerCase()) === -1)
               $(this).hide();
            else
               $(this).show();                
        });
    }); 
				$("#search2").keyup(function(){
        _this = this;
        // Show only matching TR, hide rest of them
        $.each($("#vaildBridges .t2"), function() {
            if($(this).text().toLowerCase().indexOf($(_this).val().toLowerCase()) === -1)
               $(this).hide();
            else
               $(this).show();                
        });
    }); 
				
				
			},
			error: function (req, status, err) {
				console.log('Something went wrong', status, err);
				console.log(err)

			}
		});
	});
});

