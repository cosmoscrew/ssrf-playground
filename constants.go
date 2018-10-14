package main

// URLInput is the HTML template to insert a URL and make a GET request to the same page
var URLInput = `<form action="" method="GET">
<fieldset>
  <div class="form-group">
	<div class="col-sm-5">
		<label for="staticURL">URL</label>
		<input type="text" class="form-control" name="url" id="staticURL" aria-describedby="urlHelp" placeholder="Enter a URL">
		<small id="urlHelp" class="form-text text-muted">Enter a URL to fetch content from.</small><br>
		<button type="submit" class="btn btn-primary">Submit</button>
	</div>
  </div>
</fieldset>
</form>`
