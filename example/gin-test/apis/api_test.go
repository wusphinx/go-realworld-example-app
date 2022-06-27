package apis

import (
	"bytes"
	"testing"

	"github.com/wusphinx/go-realworld-example-app/example/libs/itest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_startPage400(t *testing.T) {
	r := gin.Default()
	Register(r)
	raw := `{"address":"test"}`
	w := itest.PerformRequest(r, "POST", "/testing", bytes.NewBufferString(raw), nil)
	assert.Equal(t, w.Code, 400)
	t.Logf("resp: %s", w.Body.String())
}

func Test_startPage200(t *testing.T) {
	r := gin.Default()
	Register(r)
	raw := `{"name":"now","address":"test"}`
	w := itest.PerformRequest(r, "POST", "/testing", bytes.NewBufferString(raw), nil)
	assert.Equal(t, w.Code, 200)
	t.Logf("resp: %s", w.Body.String())
}
