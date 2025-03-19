use strict;
use warnings;
use feature 'say';

#define character sets
my $lowercase = 'abcdefghijklmnopqrstuvwxyz';
my $uppercase = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
my $numbers   = '0123456789';
my $special   = '!@#$%^&*()-_=+[]{}|;:,./?';

#get user preferences 
print "Enter password length: ";
chomp(my $length = <STDIN>);

print "Include uppercase letters? (y/n): ";
chomp(my $include_uppercase = <STDIN>);

print "Include numbers? (y/n): ";
chomp(my $include_numbers = <STDIN>);

print "Include special characters? (y/n): ";
chomp(my $include_special = <STDIN>);

#build character set
my $charset = $lowercase;
$charset .= $uppercase if $include_uppercase =~ /y/i;
$charset .= $numbers if $include_numbers =~ /y/i;
$charset .= $special if $include_special =~ /y/i;

die "No character sets selected!\n" if length ($charset) == 0;

#generate password
my $password = '';
$password .= substr($charset, int(rand(length($charset))), 1) for (1..$length);

#print generated password
say "Generated Password: $password";
