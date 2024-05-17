<!-----



Conversion time: 0.304 seconds.


Using this Markdown file:

1. Paste this output into your source file.
2. See the notes and action items below regarding this conversion run.
3. Check the rendered output (headings, lists, code blocks, tables) for proper
   formatting and use a linkchecker before you publish this page.

Conversion notes:

* Docs to Markdown version 1.0β36
* Fri May 17 2024 06:30:35 GMT-0700 (PDT)
* Source doc: Machine Update
----->



# Machine Update

<p style="text-align: right">
Brotherlogic</p>


<p style="text-align: right">
2024-05-16</p>


<p style="text-align: right">
Draft</p>



### Abstract

This document describes a process to keep machines up to date, and to continually ensure that machines are updated.


### Process

There are two things in play:



1. Machines should follow a standard set of configuration - i.e. each machine should be updated according to the machine github repo
2. Machines should be regularly updated when we can

To achieve this we add two fields to the MDB:



1. &lt;string> version
2. Int64 last_updated (timestamp)

Then we have two different ansible scripts - one which applies a given version and one which updates a given machine to be fully up to date.

We update machines after 24 hours, and we apply configuration on every update to the main database. So for machines we need to track (a) which version it’s running on and (b) when it was last updated. The ansible script then runs 5 minutes and determines (a) which machines are available and (b), if they are, what type of update is needed for them. If any hosts are found in either category, the update is applied and the MDB is updated accordingly. If an update or configuration fails then we do not update the MDB. In this manner, machines are kept up to date and new configuration is applied when the machine is available. We can also raise an issue if a machine hasn’t been updated or is behind on its configuration 


### Milestones



1. Store this document in the MDB repo
2. Update the MDB configuration to add the new fields
3. Loads the github version into the ansible playbook for a configuration update
4. Refactor the host generation script to elide empty groups (i.e. we don’t add a heading when no machines are found)
5. The host file generator for update uses the version information to filter out non-applicable machines.
6. Ansible should self add the cron job to add the update job
7. Update job should filter machines out on the basis of time
8. Ansible cron should run both update and version sync jobs on a five minute schedule