// Code generated by github.com/gobuffalo/packr. DO NOT EDIT.

package migrate

import "github.com/gobuffalo/packr"

// You can use the "packr clean" command to clean up this,
// and any other packr generated files.
func init() {
	packr.PackJSONBytes("./sql", "20180103142001_initial_schema.sql", "\"LyoKICogQ29weXJpZ2h0IDIwMTggVGhlIE5ha2FtYSBBdXRob3JzCiAqCiAqIExpY2Vuc2VkIHVuZGVyIHRoZSBBcGFjaGUgTGljZW5zZSwgVmVyc2lvbiAyLjAgKHRoZSAiTGljZW5zZSIpOwogKiB5b3UgbWF5IG5vdCB1c2UgdGhpcyBmaWxlIGV4Y2VwdCBpbiBjb21wbGlhbmNlIHdpdGggdGhlIExpY2Vuc2UuCiAqIFlvdSBtYXkgb2J0YWluIGEgY29weSBvZiB0aGUgTGljZW5zZSBhdAogKgogKiBodHRwOi8vd3d3LmFwYWNoZS5vcmcvbGljZW5zZXMvTElDRU5TRS0yLjAKICoKICogVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQogKiBkaXN0cmlidXRlZCB1bmRlciB0aGUgTGljZW5zZSBpcyBkaXN0cmlidXRlZCBvbiBhbiAiQVMgSVMiIEJBU0lTLAogKiBXSVRIT1VUIFdBUlJBTlRJRVMgT1IgQ09ORElUSU9OUyBPRiBBTlkgS0lORCwgZWl0aGVyIGV4cHJlc3Mgb3IgaW1wbGllZC4KICogU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZAogKiBsaW1pdGF0aW9ucyB1bmRlciB0aGUgTGljZW5zZS4KICovCgotLSArbWlncmF0ZSBVcApDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyB1c2VycyAoCiAgICBQUklNQVJZIEtFWSAoaWQpLAoKICAgIGlkICAgICAgICAgICAgVVVJRCAgICAgICAgICBOT1QgTlVMTCwKICAgIHVzZXJuYW1lICAgICAgVkFSQ0hBUigxMjgpICBDT05TVFJBSU5UIHVzZXJzX3VzZXJuYW1lX2tleSBVTklRVUUgTk9UIE5VTEwsCiAgICBkaXNwbGF5X25hbWUgIFZBUkNIQVIoMjU1KSwKICAgIGF2YXRhcl91cmwgICAgVkFSQ0hBUig1MTIpLAogICAgLS0gaHR0cHM6Ly90b29scy5pZXRmLm9yZy9odG1sL2JjcDQ3CiAgICBsYW5nX3RhZyAgICAgIFZBUkNIQVIoMTgpICAgREVGQVVMVCAnZW4nLAogICAgbG9jYXRpb24gICAgICBWQVJDSEFSKDI1NSksIC0tIGUuZy4gIlNhbiBGcmFuY2lzY28sIENBIgogICAgdGltZXpvbmUgICAgICBWQVJDSEFSKDI1NSksIC0tIGUuZy4gIlBhY2lmaWMgVGltZSAoVVMgJiBDYW5hZGEpIgogICAgbWV0YWRhdGEgICAgICBKU09OQiAgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwKICAgIHdhbGxldCAgICAgICAgSlNPTkIgICAgICAgICBERUZBVUxUICd7fScgTk9UIE5VTEwsCiAgICBlbWFpbCAgICAgICAgIFZBUkNIQVIoMjU1KSAgVU5JUVVFLAogICAgcGFzc3dvcmQgICAgICBCWVRFQSAgICAgICAgIENIRUNLIChsZW5ndGgocGFzc3dvcmQpIDwgMzIwMDApLAogICAgZmFjZWJvb2tfaWQgICBWQVJDSEFSKDEyOCkgIFVOSVFVRSwKICAgIGdvb2dsZV9pZCAgICAgVkFSQ0hBUigxMjgpICBVTklRVUUsCiAgICBnYW1lY2VudGVyX2lkIFZBUkNIQVIoMTI4KSAgVU5JUVVFLAogICAgc3RlYW1faWQgICAgICBWQVJDSEFSKDEyOCkgIFVOSVFVRSwKICAgIGN1c3RvbV9pZCAgICAgVkFSQ0hBUigxMjgpICBVTklRVUUsCiAgICBlZGdlX2NvdW50ICAgIElOVCAgICAgICAgICAgREVGQVVMVCAwIENIRUNLIChlZGdlX2NvdW50ID49IDApIE5PVCBOVUxMLAogICAgY3JlYXRlX3RpbWUgICBUSU1FU1RBTVBUWiAgIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCiAgICB1cGRhdGVfdGltZSAgIFRJTUVTVEFNUFRaICAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwKICAgIHZlcmlmeV90aW1lICAgVElNRVNUQU1QVFogICBERUZBVUxUICcxOTcwLTAxLTAxIDAwOjAwOjAwJyBOT1QgTlVMTCwKICAgIGRpc2FibGVfdGltZSAgVElNRVNUQU1QVFogICBERUZBVUxUICcxOTcwLTAxLTAxIDAwOjAwOjAwJyBOT1QgTlVMTAopOwoKLS0gU2V0dXAgU3lzdGVtIHVzZXIuCklOU0VSVCBJTlRPIHVzZXJzIChpZCwgdXNlcm5hbWUpCiAgICBWQUxVRVMgKCcwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAnLCAnJykKICAgIE9OIENPTkZMSUNUKGlkKSBETyBOT1RISU5HOwoKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgdXNlcl9kZXZpY2UgKAogICAgUFJJTUFSWSBLRVkgKGlkKSwKICAgIEZPUkVJR04gS0VZICh1c2VyX2lkKSBSRUZFUkVOQ0VTIHVzZXJzIChpZCkgT04gREVMRVRFIENBU0NBREUsCgogICAgaWQgICAgICBWQVJDSEFSKDEyOCkgTk9UIE5VTEwsCiAgICB1c2VyX2lkIFVVSUQgICAgICAgICBOT1QgTlVMTAopOwoKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgdXNlcl9lZGdlICgKICAgIFBSSU1BUlkgS0VZIChzb3VyY2VfaWQsIHN0YXRlLCBwb3NpdGlvbiksCiAgICBGT1JFSUdOIEtFWSAoc291cmNlX2lkKSAgICAgIFJFRkVSRU5DRVMgdXNlcnMgKGlkKSBPTiBERUxFVEUgQ0FTQ0FERSwKICAgIEZPUkVJR04gS0VZIChkZXN0aW5hdGlvbl9pZCkgUkVGRVJFTkNFUyB1c2VycyAoaWQpIE9OIERFTEVURSBDQVNDQURFLAoKICAgIHNvdXJjZV9pZCAgICAgIFVVSUQgICAgICAgIE5PVCBOVUxMLAogICAgcG9zaXRpb24gICAgICAgQklHSU5UICAgICAgTk9UIE5VTEwsIC0tIFVzZWQgZm9yIHNvcnQgb3JkZXIgb24gcm93cy4KICAgIHVwZGF0ZV90aW1lICAgIFRJTUVTVEFNUFRaIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCiAgICBkZXN0aW5hdGlvbl9pZCBVVUlEICAgICAgICBOT1QgTlVMTCwKICAgIHN0YXRlICAgICAgICAgIFNNQUxMSU5UICAgIERFRkFVTFQgMCBOT1QgTlVMTCwgLS0gZnJpZW5kKDApLCBpbnZpdGUoMSksIGludml0ZWQoMiksIGJsb2NrZWQoMyksIGRlbGV0ZWQoNCksIGFyY2hpdmVkKDUpCgogICAgVU5JUVVFIChzb3VyY2VfaWQsIGRlc3RpbmF0aW9uX2lkKQopOwoKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgbm90aWZpY2F0aW9uICgKICAgIC0tIEZJWE1FOiBjb2Nrcm9hY2gncyBhbmFseXNlciBpcyBub3QgY2xldmVyIGVub3VnaCB3aGVuIGNyZWF0ZV90aW1lIGhhcyBERVNDIG1vZGUgb24gdGhlIGluZGV4LgogICAgUFJJTUFSWSBLRVkgKHVzZXJfaWQsIGNyZWF0ZV90aW1lLCBpZCksCiAgICBGT1JFSUdOIEtFWSAodXNlcl9pZCkgUkVGRVJFTkNFUyB1c2VycyAoaWQpIE9OIERFTEVURSBDQVNDQURFLAoKICAgIGlkICAgICAgICAgIFVVSUQgICAgICAgICBDT05TVFJBSU5UIG5vdGlmaWNhdGlvbl9pZF9rZXkgVU5JUVVFIE5PVCBOVUxMLAogICAgdXNlcl9pZCAgICAgVVVJRCAgICAgICAgIE5PVCBOVUxMLAogICAgc3ViamVjdCAgICAgVkFSQ0hBUigyNTUpIE5PVCBOVUxMLAogICAgY29udGVudCAgICAgSlNPTkIgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwKICAgIGNvZGUgICAgICAgIFNNQUxMSU5UICAgICBOT1QgTlVMTCwgLS0gTmVnYXRpdmUgdmFsdWVzIGFyZSBzeXN0ZW0gcmVzZXJ2ZWQuCiAgICBzZW5kZXJfaWQgICBVVUlEICAgICAgICAgTk9UIE5VTEwsCiAgICBjcmVhdGVfdGltZSBUSU1FU1RBTVBUWiAgREVGQVVMVCBub3coKSBOT1QgTlVMTAopOwoKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgc3RvcmFnZSAoCiAgICBQUklNQVJZIEtFWSAoY29sbGVjdGlvbiwgcmVhZCwga2V5LCB1c2VyX2lkKSwKICAgIEZPUkVJR04gS0VZICh1c2VyX2lkKSBSRUZFUkVOQ0VTIHVzZXJzIChpZCkgT04gREVMRVRFIENBU0NBREUsCgogICAgY29sbGVjdGlvbiAgVkFSQ0hBUigxMjgpIE5PVCBOVUxMLAogICAga2V5ICAgICAgICAgVkFSQ0hBUigxMjgpIE5PVCBOVUxMLAogICAgdXNlcl9pZCAgICAgVVVJRCAgICAgICAgIE5PVCBOVUxMLAogICAgdmFsdWUgICAgICAgSlNPTkIgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwKICAgIHZlcnNpb24gICAgIFZBUkNIQVIoMzIpICBOT1QgTlVMTCwgLS0gbWQ1IGhhc2ggb2YgdmFsdWUgb2JqZWN0LgogICAgcmVhZCAgICAgICAgU01BTExJTlQgICAgIERFRkFVTFQgMSBDSEVDSyAocmVhZCA+PSAwKSBOT1QgTlVMTCwKICAgIHdyaXRlICAgICAgIFNNQUxMSU5UICAgICBERUZBVUxUIDEgQ0hFQ0sgKHdyaXRlID49IDApIE5PVCBOVUxMLAogICAgY3JlYXRlX3RpbWUgVElNRVNUQU1QVFogIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCiAgICB1cGRhdGVfdGltZSBUSU1FU1RBTVBUWiAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwKCiAgICBVTklRVUUgKGNvbGxlY3Rpb24sIGtleSwgdXNlcl9pZCkKKTsKQ1JFQVRFIElOREVYIElGIE5PVCBFWElTVFMgY29sbGVjdGlvbl9yZWFkX3VzZXJfaWRfa2V5X2lkeCBPTiBzdG9yYWdlIChjb2xsZWN0aW9uLCByZWFkLCB1c2VyX2lkLCBrZXkpOwpDUkVBVEUgSU5ERVggSUYgTk9UIEVYSVNUUyB2YWx1ZV9naW5pZHggT04gc3RvcmFnZSBVU0lORyBHSU4gKHZhbHVlKTsKCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIG1lc3NhZ2UgKAogIFBSSU1BUlkgS0VZIChzdHJlYW1fbW9kZSwgc3RyZWFtX3N1YmplY3QsIHN0cmVhbV9kZXNjcmlwdG9yLCBzdHJlYW1fbGFiZWwsIGNyZWF0ZV90aW1lLCBpZCksCiAgRk9SRUlHTiBLRVkgKHNlbmRlcl9pZCkgUkVGRVJFTkNFUyB1c2VycyAoaWQpIE9OIERFTEVURSBDQVNDQURFLAoKICBpZCAgICAgICAgICAgICAgICBVVUlEICAgICAgICAgVU5JUVVFIE5PVCBOVUxMLAogIC0tIGNoYXQoMCksIGNoYXRfdXBkYXRlKDEpLCBjaGF0X3JlbW92ZSgyKSwgZ3JvdXBfam9pbigzKSwgZ3JvdXBfYWRkKDQpLCBncm91cF9sZWF2ZSg1KSwgZ3JvdXBfa2ljayg2KSwgZ3JvdXBfcHJvbW90ZWQoNykKICBjb2RlICAgICAgICAgICAgICBTTUFMTElOVCAgICAgREVGQVVMVCAwIE5PVCBOVUxMLAogIHNlbmRlcl9pZCAgICAgICAgIFVVSUQgICAgICAgICBOT1QgTlVMTCwKICB1c2VybmFtZSAgICAgICAgICBWQVJDSEFSKDEyOCkgTk9UIE5VTEwsCiAgc3RyZWFtX21vZGUgICAgICAgU01BTExJTlQgICAgIE5PVCBOVUxMLAogIHN0cmVhbV9zdWJqZWN0ICAgIFVVSUQgICAgICAgICBOT1QgTlVMTCwKICBzdHJlYW1fZGVzY3JpcHRvciBVVUlEICAgICAgICAgTk9UIE5VTEwsCiAgc3RyZWFtX2xhYmVsICAgICAgVkFSQ0hBUigxMjgpIE5PVCBOVUxMLAogIGNvbnRlbnQgICAgICAgICAgIEpTT05CICAgICAgICBERUZBVUxUICd7fScgTk9UIE5VTEwsCiAgY3JlYXRlX3RpbWUgICAgICAgVElNRVNUQU1QVFogIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCiAgdXBkYXRlX3RpbWUgICAgICAgVElNRVNUQU1QVFogIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCgogIFVOSVFVRSAoc2VuZGVyX2lkLCBpZCkKKTsKCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIGxlYWRlcmJvYXJkICgKICBQUklNQVJZIEtFWSAoaWQpLAoKICBpZCAgICAgICAgICAgICBWQVJDSEFSKDEyOCkgTk9UIE5VTEwsCiAgYXV0aG9yaXRhdGl2ZSAgQk9PTEVBTiAgICAgIERFRkFVTFQgRkFMU0UsCiAgc29ydF9vcmRlciAgICAgU01BTExJTlQgICAgIERFRkFVTFQgMSBOT1QgTlVMTCwgLS0gYXNjKDApLCBkZXNjKDEpCiAgb3BlcmF0b3IgICAgICAgU01BTExJTlQgICAgIERFRkFVTFQgMCBOT1QgTlVMTCwgLS0gYmVzdCgwKSwgc2V0KDEpLCBpbmNyZW1lbnQoMiksIGRlY3JlbWVudCgzKQogIHJlc2V0X3NjaGVkdWxlIFZBUkNIQVIoNjQpLCAtLSBlLmcuIGNyb24gZm9ybWF0OiAiKiAqICogKiAqICogKiIKICBtZXRhZGF0YSAgICAgICBKU09OQiAgICAgICAgREVGQVVMVCAne30nIE5PVCBOVUxMLAogIGNyZWF0ZV90aW1lICAgIFRJTUVTVEFNUFRaICBERUZBVUxUIG5vdygpIE5PVCBOVUxMCik7CgpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyBsZWFkZXJib2FyZF9yZWNvcmQgKAogIFBSSU1BUlkgS0VZIChsZWFkZXJib2FyZF9pZCwgZXhwaXJ5X3RpbWUsIHNjb3JlLCBzdWJzY29yZSwgb3duZXJfaWQpLAogIEZPUkVJR04gS0VZIChsZWFkZXJib2FyZF9pZCkgUkVGRVJFTkNFUyBsZWFkZXJib2FyZCAoaWQpIE9OIERFTEVURSBDQVNDQURFLAoKICBsZWFkZXJib2FyZF9pZCBWQVJDSEFSKDEyOCkgIE5PVCBOVUxMLAogIG93bmVyX2lkICAgICAgIFVVSUQgICAgICAgICAgTk9UIE5VTEwsCiAgdXNlcm5hbWUgICAgICAgVkFSQ0hBUigxMjgpLAogIHNjb3JlICAgICAgICAgIEJJR0lOVCAgICAgICAgREVGQVVMVCAwIENIRUNLIChzY29yZSA+PSAwKSBOT1QgTlVMTCwKICBzdWJzY29yZSAgICAgICBCSUdJTlQgICAgICAgIERFRkFVTFQgMCBDSEVDSyAoc3Vic2NvcmUgPj0gMCkgTk9UIE5VTEwsCiAgbnVtX3Njb3JlICAgICAgSU5UICAgICAgICAgICBERUZBVUxUIDEgQ0hFQ0sgKG51bV9zY29yZSA+PSAwKSBOT1QgTlVMTCwKICBtZXRhZGF0YSAgICAgICBKU09OQiAgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwKICBjcmVhdGVfdGltZSAgICBUSU1FU1RBTVBUWiAgIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCiAgdXBkYXRlX3RpbWUgICAgVElNRVNUQU1QVFogICBERUZBVUxUIG5vdygpIE5PVCBOVUxMLAogIGV4cGlyeV90aW1lICAgIFRJTUVTVEFNUFRaICAgREVGQVVMVCAnMTk3MC0wMS0wMSAwMDowMDowMCcgTk9UIE5VTEwsCgogIFVOSVFVRSAob3duZXJfaWQsIGxlYWRlcmJvYXJkX2lkLCBleHBpcnlfdGltZSkKKTsKCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIHdhbGxldF9sZWRnZXIgKAogIFBSSU1BUlkgS0VZICh1c2VyX2lkLCBjcmVhdGVfdGltZSwgaWQpLAogIEZPUkVJR04gS0VZICh1c2VyX2lkKSBSRUZFUkVOQ0VTIHVzZXJzIChpZCkgT04gREVMRVRFIENBU0NBREUsCgogIGlkICAgICAgICAgIFVVSUQgICAgICAgIFVOSVFVRSBOT1QgTlVMTCwKICB1c2VyX2lkICAgICBVVUlEICAgICAgICBOT1QgTlVMTCwKICBjaGFuZ2VzZXQgICBKU09OQiAgICAgICBOT1QgTlVMTCwKICBtZXRhZGF0YSAgICBKU09OQiAgICAgICBOT1QgTlVMTCwKICBjcmVhdGVfdGltZSBUSU1FU1RBTVBUWiBERUZBVUxUIG5vdygpIE5PVCBOVUxMLAogIHVwZGF0ZV90aW1lIFRJTUVTVEFNUFRaIERFRkFVTFQgbm93KCkgTk9UIE5VTEwKKTsKCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIHVzZXJfdG9tYnN0b25lICgKICBQUklNQVJZIEtFWSAoY3JlYXRlX3RpbWUsIHVzZXJfaWQpLAoKICB1c2VyX2lkICAgICAgICBVVUlEICAgICAgICBVTklRVUUgTk9UIE5VTEwsCiAgY3JlYXRlX3RpbWUgICAgVElNRVNUQU1QVFogREVGQVVMVCBub3coKSBOT1QgTlVMTAopOwoKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgZ3JvdXBzICgKICBQUklNQVJZIEtFWSAoZGlzYWJsZV90aW1lLCBsYW5nX3RhZywgZWRnZV9jb3VudCwgaWQpLAoKICBpZCAgICAgICAgICAgVVVJRCAgICAgICAgICBVTklRVUUgTk9UIE5VTEwsCiAgY3JlYXRvcl9pZCAgIFVVSUQgICAgICAgICAgTk9UIE5VTEwsCiAgbmFtZSAgICAgICAgIFZBUkNIQVIoMjU1KSAgQ09OU1RSQUlOVCBncm91cHNfbmFtZV9rZXkgVU5JUVVFIE5PVCBOVUxMLAogIGRlc2NyaXB0aW9uICBWQVJDSEFSKDI1NSksCiAgYXZhdGFyX3VybCAgIFZBUkNIQVIoNTEyKSwKICAtLSBodHRwczovL3Rvb2xzLmlldGYub3JnL2h0bWwvYmNwNDcKICBsYW5nX3RhZyAgICAgVkFSQ0hBUigxOCkgICBERUZBVUxUICdlbicsCiAgbWV0YWRhdGEgICAgIEpTT05CICAgICAgICAgREVGQVVMVCAne30nIE5PVCBOVUxMLAogIHN0YXRlICAgICAgICBTTUFMTElOVCAgICAgIERFRkFVTFQgMCBDSEVDSyAoc3RhdGUgPj0gMCkgTk9UIE5VTEwsIC0tIG9wZW4oMCksIGNsb3NlZCgxKQogIGVkZ2VfY291bnQgICBJTlQgICAgICAgICAgIERFRkFVTFQgMCBDSEVDSyAoZWRnZV9jb3VudCA+PSAxIEFORCBlZGdlX2NvdW50IDw9IG1heF9jb3VudCkgTk9UIE5VTEwsCiAgbWF4X2NvdW50ICAgIElOVCAgICAgICAgICAgREVGQVVMVCAxMDAgQ0hFQ0sgKG1heF9jb3VudCA+PSAxKSBOT1QgTlVMTCwKICBjcmVhdGVfdGltZSAgVElNRVNUQU1QVFogICBERUZBVUxUIG5vdygpIE5PVCBOVUxMLAogIHVwZGF0ZV90aW1lICBUSU1FU1RBTVBUWiAgIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsCiAgZGlzYWJsZV90aW1lIFRJTUVTVEFNUFRaICAgREVGQVVMVCAnMTk3MC0wMS0wMSAwMDowMDowMCcgTk9UIE5VTEwKKTsKQ1JFQVRFIElOREVYIElGIE5PVCBFWElTVFMgZWRnZV9jb3VudF91cGRhdGVfdGltZV9pZF9pZHggT04gZ3JvdXBzIChkaXNhYmxlX3RpbWUsIGVkZ2VfY291bnQsIHVwZGF0ZV90aW1lLCBpZCk7CkNSRUFURSBJTkRFWCBJRiBOT1QgRVhJU1RTIHVwZGF0ZV90aW1lX2VkZ2VfY291bnRfaWRfaWR4IE9OIGdyb3VwcyAoZGlzYWJsZV90aW1lLCB1cGRhdGVfdGltZSwgZWRnZV9jb3VudCwgaWQpOwoKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgZ3JvdXBfZWRnZSAoCiAgUFJJTUFSWSBLRVkgKHNvdXJjZV9pZCwgc3RhdGUsIHBvc2l0aW9uKSwKCiAgc291cmNlX2lkICAgICAgVVVJRCAgICAgICAgTk9UIE5VTEwsCiAgcG9zaXRpb24gICAgICAgQklHSU5UICAgICAgTk9UIE5VTEwsIC0tIFVzZWQgZm9yIHNvcnQgb3JkZXIgb24gcm93cy4KICB1cGRhdGVfdGltZSAgICBUSU1FU1RBTVBUWiBERUZBVUxUIG5vdygpIE5PVCBOVUxMLAogIGRlc3RpbmF0aW9uX2lkIFVVSUQgICAgICAgIE5PVCBOVUxMLAogIHN0YXRlICAgICAgICAgIFNNQUxMSU5UICAgIERFRkFVTFQgMCBOT1QgTlVMTCwgLS0gc3VwZXJhZG1pbigwKSwgYWRtaW4oMSksIG1lbWJlcigyKSwgam9pbl9yZXF1ZXN0KDMpLCBhcmNoaXZlZCg0KQoKICBVTklRVUUgKHNvdXJjZV9pZCwgZGVzdGluYXRpb25faWQpCik7CgotLSArbWlncmF0ZSBEb3duCkRST1AgVEFCTEUgSUYgRVhJU1RTIGdyb3VwX2VkZ2U7CkRST1AgVEFCTEUgSUYgRVhJU1RTIGdyb3VwczsKRFJPUCBUQUJMRSBJRiBFWElTVFMgdXNlcl90b21ic3RvbmU7CkRST1AgVEFCTEUgSUYgRVhJU1RTIHdhbGxldF9sZWRnZXI7CkRST1AgVEFCTEUgSUYgRVhJU1RTIGxlYWRlcmJvYXJkX3JlY29yZDsKRFJPUCBUQUJMRSBJRiBFWElTVFMgbGVhZGVyYm9hcmQ7CkRST1AgVEFCTEUgSUYgRVhJU1RTIG1lc3NhZ2U7CkRST1AgVEFCTEUgSUYgRVhJU1RTIHN0b3JhZ2U7CkRST1AgVEFCTEUgSUYgRVhJU1RTIG5vdGlmaWNhdGlvbjsKRFJPUCBUQUJMRSBJRiBFWElTVFMgdXNlcl9lZGdlOwpEUk9QIFRBQkxFIElGIEVYSVNUUyB1c2VyX2RldmljZTsKRFJPUCBUQUJMRSBJRiBFWElTVFMgdXNlcnM7Cg==\"")
	packr.PackJSONBytes("./sql", "20180805174141-tournaments.sql", "\"LyoKICogQ29weXJpZ2h0IDIwMTggVGhlIE5ha2FtYSBBdXRob3JzCiAqCiAqIExpY2Vuc2VkIHVuZGVyIHRoZSBBcGFjaGUgTGljZW5zZSwgVmVyc2lvbiAyLjAgKHRoZSAiTGljZW5zZSIpOwogKiB5b3UgbWF5IG5vdCB1c2UgdGhpcyBmaWxlIGV4Y2VwdCBpbiBjb21wbGlhbmNlIHdpdGggdGhlIExpY2Vuc2UuCiAqIFlvdSBtYXkgb2J0YWluIGEgY29weSBvZiB0aGUgTGljZW5zZSBhdAogKgogKiBodHRwOi8vd3d3LmFwYWNoZS5vcmcvbGljZW5zZXMvTElDRU5TRS0yLjAKICoKICogVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQogKiBkaXN0cmlidXRlZCB1bmRlciB0aGUgTGljZW5zZSBpcyBkaXN0cmlidXRlZCBvbiBhbiAiQVMgSVMiIEJBU0lTLAogKiBXSVRIT1VUIFdBUlJBTlRJRVMgT1IgQ09ORElUSU9OUyBPRiBBTlkgS0lORCwgZWl0aGVyIGV4cHJlc3Mgb3IgaW1wbGllZC4KICogU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZAogKiBsaW1pdGF0aW9ucyB1bmRlciB0aGUgTGljZW5zZS4KICovCgotLSBOT1RFOiBUaGlzIG1pZ3JhdGlvbiBtYW51YWxseSBjb21taXRzIGluIHNlcGFyYXRlIHRyYW5zYWN0aW9ucyB0byBlbnN1cmUKLS0gdGhlIHNjaGVtYSB1cGRhdGVzIGFyZSBzZXF1ZW5jZWQgYmVjYXVzZSBjb2Nrcm9hY2hkYiBkb2VzIG5vdCBzdXBwb3J0Ci0tIGFkZGluZyBDSEVDSyBjb25zdHJhaW50cyB2aWEgIkFMVEVSIFRBQkxFIC4uLiBBREQgQ09MVU1OIiBzdGF0ZW1lbnRzLgoKLS0gK21pZ3JhdGUgVXAgbm90cmFuc2FjdGlvbgpCRUdJTjsKQUxURVIgVEFCTEUgbGVhZGVyYm9hcmQKICBBREQgQ09MVU1OIGNhdGVnb3J5ICAgICAgU01BTExJTlQgICAgIERFRkFVTFQgMCAvKkNIRUNLIChjYXRlZ29yeSA+PSAwKSovIE5PVCBOVUxMLAogIEFERCBDT0xVTU4gZGVzY3JpcHRpb24gICBWQVJDSEFSKDI1NSkgREVGQVVMVCAnJyBOT1QgTlVMTCwKICBBREQgQ09MVU1OIGR1cmF0aW9uICAgICAgSU5UICAgICAgICAgIERFRkFVTFQgODY0MDAgLypDSEVDSyAoZHVyYXRpb24gPiAwKSovIE5PVCBOVUxMLCAtLSBpbiBzZWNvbmRzLgogIEFERCBDT0xVTU4gZW5kX3RpbWUgICAgICBUSU1FU1RBTVBUWiwKICBBREQgQ09MVU1OIG1heF9zaXplICAgICAgSU5UICAgICAgICAgIERFRkFVTFQgMTAwMDAwMDAwIC8qQ0hFQ0sgKG1heF9zaXplID4gMCkqLyBOT1QgTlVMTCwKICBBREQgQ09MVU1OIG1heF9udW1fc2NvcmUgSU5UICAgICAgICAgIERFRkFVTFQgMTAwMDAwMCAvKkNIRUNLIChtYXhfbnVtX3Njb3JlID4gMCkqLyBOT1QgTlVMTCwgLS0gbWF4IGFsbG93ZWQgc2NvcmUgYXR0ZW1wdHMuCiAgQUREIENPTFVNTiBuYW1lICAgICAgICAgIFZBUkNIQVIoMjU1KSBERUZBVUxUICcnIE5PVCBOVUxMLAogIEFERCBDT0xVTU4gc2l6ZSAgICAgICAgICBJTlQgICAgICAgICAgREVGQVVMVCAwIE5PVCBOVUxMLAogIEFERCBDT0xVTU4gc3RhcnRfdGltZSAgICBUSU1FU1RBTVBUWjsKCkFMVEVSIFRBQkxFIGxlYWRlcmJvYXJkX3JlY29yZAogIEFERCBDT0xVTU4gbWF4X251bV9zY29yZSBJTlQgREVGQVVMVCAxMDAwMDAwIC8qQ0hFQ0sgKG1heF9udW1fc2NvcmUgPiAwKSovIE5PVCBOVUxMOwpDT01NSVQ7CgpCRUdJTjsKQUxURVIgVEFCTEUgbGVhZGVyYm9hcmQKICBBREQgQ09OU1RSQUlOVCBjaGVja19jYXRlZ29yeSBDSEVDSyAoY2F0ZWdvcnkgPj0gMCksCiAgQUREIENPTlNUUkFJTlQgY2hlY2tfZHVyYXRpb24gQ0hFQ0sgKGR1cmF0aW9uID4gMCksCiAgQUREIENPTlNUUkFJTlQgY2hlY2tfbWF4X3NpemUgQ0hFQ0sgKG1heF9zaXplID4gMCksCiAgQUREIENPTlNUUkFJTlQgY2hlY2tfbWF4X251bV9zY29yZSBDSEVDSyAobWF4X251bV9zY29yZSA+IDApLAogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfY2F0ZWdvcnksCiAgVkFMSURBVEUgQ09OU1RSQUlOVCBjaGVja19kdXJhdGlvbiwKICBWQUxJREFURSBDT05TVFJBSU5UIGNoZWNrX21heF9zaXplLAogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfbWF4X251bV9zY29yZTsKCkFMVEVSIFRBQkxFIGxlYWRlcmJvYXJkX3JlY29yZAogIEFERCBDT05TVFJBSU5UIGNoZWNrX21heF9udW1fc2NvcmUgQ0hFQ0sgKG1heF9udW1fc2NvcmUgPiAwKSwKICBWQUxJREFURSBDT05TVFJBSU5UIGNoZWNrX21heF9udW1fc2NvcmU7CkNPTU1JVDsKCi0tICttaWdyYXRlIERvd24KQUxURVIgVEFCTEUgSUYgRVhJU1RTIGxlYWRlcmJvYXJkCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIGNhdGVnb3J5LAogIERST1AgQ09MVU1OIElGIEVYSVNUUyBkZXNjcmlwdGlvbiwKICBEUk9QIENPTFVNTiBJRiBFWElTVFMgZHVyYXRpb24sCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIGVuZF90aW1lLAogIERST1AgQ09MVU1OIElGIEVYSVNUUyBtYXhfc2l6ZSwKICBEUk9QIENPTFVNTiBJRiBFWElTVFMgbWF4X251bV9zY29yZSwKICBEUk9QIENPTFVNTiBJRiBFWElTVFMgbmFtZSwKICBEUk9QIENPTFVNTiBJRiBFWElTVFMgc2l6ZSwKICBEUk9QIENPTFVNTiBJRiBFWElTVFMgc3RhcnRfdGltZTsKCkFMVEVSIFRBQkxFIElGIEVYSVNUUyBsZWFkZXJib2FyZF9yZWNvcmQKICBEUk9QIENPTFVNTiBJRiBFWElTVFMgbWF4X251bV9zY29yZTsK\"")
}
