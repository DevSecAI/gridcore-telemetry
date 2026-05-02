// GRID-SAST-001: command injection — site argument is interpolated unquoted.
package scada

import (
	"os/exec"
)

// RunDiagnostic ssh's into a substation gateway and runs a diagnostic.
// `site` arrives from the API and reaches the shell directly via /bin/sh -c.
func RunDiagnostic(site string) ([]byte, error) {
	// GRID-SAST-001
	cmd := exec.Command("/bin/sh", "-c", "ssh -o StrictHostKeyChecking=no operator@"+site+" /opt/grid/diag.sh")
	return cmd.CombinedOutput()
}
