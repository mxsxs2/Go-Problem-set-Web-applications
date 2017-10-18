 //Wait until the document is loaded
 jQuery(document).ready(function(){
    //Set the focus on the field
    jQuery("input[name='guess']").focus();
    //Submit the form on enter 
    jQuery("#guessForm").on("keypress",function(){
      //If the key is enter
      if (e.which == 13) {
        //Submit the form
        jQuery("#guessForm").submit();
        return false;
      }
    })
});