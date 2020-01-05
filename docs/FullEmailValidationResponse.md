# FullEmailValidationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidAddress** | **bool** | True if the email address is valid overall, false otherwise | [optional] [default to null]
**MailServerUsedForValidation** | **string** | Email server connected to for verification | [optional] [default to null]
**ValidSyntax** | **bool** | True if the syntax of the email address is valid, false otherwise.  This is one component of ValidAddress, but not the only one. | [optional] [default to null]
**ValidDomain** | **bool** | True if the domain name of the email address is valid, false otherwise.  This is one component of ValidAddress, but not the only one. | [optional] [default to null]
**ValidSMTP** | **bool** | True if the email address was verified by the remote server, false otherwise.  This is one component of ValidAddress, but not the only one. | [optional] [default to null]
**IsCatchallDomain** | **bool** | True if the domain is a catch-all domain name, false otherwise.  Catch-all domain names, while rare, always accept inbound email to ensure they do not lose any potentially useful emails.  Catch-all domain names can occassionally be configured to first accept and store all inbound email, but then later send a bounce email back to the sender after a delayed period of time. | [optional] [default to null]
**Domain** | **string** | Domain name of the email address | [optional] [default to null]
**IsFreeEmailProvider** | **bool** | True if the email domain name is a free provider (typically a free to sign up web email provider for consumers / personal use), false otherwise. | [optional] [default to null]
**IsDisposable** | **bool** | True if the email address is a disposable email address, false otherwise; these disposable providers are not typically used to receive email and so will have a low likelihood of opening mail sent there. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


