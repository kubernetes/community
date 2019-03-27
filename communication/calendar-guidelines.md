Project meetings are a life line of the Kubernetes project but calendaring is hard. Use this guide to help you navigate though the trickiness of calendars and learn from our fails.
PR in your favorite tip that can help others or if you have an example other than gmail.

### "I'm a chair for a SIG or WG and need to set up a meeting":
//This may change with the addition of a gsuite but this is the current best state   
*This calendar creation process will allow all of your leads to edit SIG/WG Meetings.*   

1- Use a poll service like doodle.com that will help you get a good pulse on your community and when they can meet  
2- Create a new shared calendar in the meantime as 'SIG Foo Shared Calendar'   
  This is important as we all change jobs, email addresses, and take breaks from the project. It will allow you to transfer the ownership to the shared calendar and then the rest of your team can edit it at anytime. [example of a shared calendar with google calendars: https://support.google.com/calendar/answer/37095?hl=en]   
3- Access permissions and sharing:
* Make all event details publicly accessible. Do this from an account that won't have problems with sharing and posting information publicly. This is important and you should test this first if you are not using a personal account like gmail. //TODO add a pic
* Share it with full rights ("make changes and manage sharing” on gmail) to: your SIG/WG lead mailing list and community@kubernetes.io. With great power comes great responsibility, let your other chairs know they can accidentally delete a calendar if they are trying to delete it from theirs.
* Lastly, share with view permissions only (“see all event details”) to: your SIG/WG mailing list   

4- Once you have a time cadence settled from your members, create a calendar invite with the shared calendar as the owner. //TODO add a pic  
5- Name it “SIG/WG Foo [Time Cadence ex: Biweekly] Meetings”  
6- Sharing: Public (note: most gmail will have a 'default visibility' setting that automatically is turned on. Default visibility is usually not public and will need to manually scroll to public)  
7- Include your meeting notes, zoom information, and any other pertinent information that you want your SIG/WG to know.   
8- Invite your SIG/WG mailing list and cgnt364vd8s86hr2phapfjc6uk@group.calendar.google.com (Why this weird address? This is a public calendar that will be used to populate calendars on various sites)  
/end

### "I'm a chair and the person that created the meeting is either no longer with the project or no longer at employer that holds the invite"
If you have a shared calendar with edit rights to other chairs, leads, etc., they can edit the invite and migate the situation. Also check with folks on the community@kubernetes.io team.  

If there is no shared calendar and still one owner, ask the person to transfer it to a shared calendar or you'll need to recreate one.
Best advice here is to recreate one. It won't hurt to recreate a meeting invite every few months anyway to refresh the group.

### "I'm a contributor and want to see one of/all of the SIG calendar(s)."
* All of the SIGs and WGs have meeting agendas with detailed information at the top. You can get this information from the SIG/WG list. Join their mailing list for the most up to date calendar invites. Chairs will always invite the entire mailing list to events.
* To see all of the meetings on one calendar: https://calendar.google.com/calendar/embed?src=cgnt364vd8s86hr2phapfjc6uk%40group.calendar.google.com&ctz=America%2FLos_Angeles

## Permissions Tips
#### If you are creating calendar events:
Make sure your work account doesn't have restrictions for public viewing of calendar invites you create. Test this with other contributors before sending it to mailing lists if you are unsure. This would be for both the calendar entry itself and the shared calendar if you are the chair creating it.
If this is the case, use a personal account (ex: gmail).

#### If you are viewing calendar events:
TODO

## Misc Tips
Don't copy calendars if you can help it. Copying the calendar onto your calendar will prevent you from getting updates like a canceled meeting.  
Always join a SIG/WG mailing list thats of interest and our main contributor list - kubernetes-dev@googlegroups.com. Accept the invite from the sender and you'll have the updates.   

If a chair is offboarding, ask them to transfer the ownership so there isn't a ghost calendar invite on your members calendar.   

//TODO - tip about timezones
