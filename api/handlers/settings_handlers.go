package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pastemyst-api/config"
	"pastemyst-api/db"
	"pastemyst-api/logging"
	"pastemyst-api/models"
	"pastemyst-api/utils"
	"pastemyst-api/validation"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// Sets the user's avatar.
//
// PATCH /api/v3/settings/avatar
func PatchAvatarHandler(ctx echo.Context) error {
	user, hasUser := ctx.Get("user").(models.User)

	if !hasUser {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing the avatar file.")
	}

	src, err := file.Open()
	if err != nil {
		logging.Logger.Errorf("Failed to open the uploaded file: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer src.Close()

	if !validation.IsValidImage(src) {
		return echo.NewHTTPError(http.StatusBadRequest, "The provided file is not a valid image file.")
	}

	avatarId := utils.RandomIdWhile(func(s string) bool {
		files, _ := os.ReadDir("./assets/avatars/")

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if s == strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) {
				return true
			}
		}

		return false
	})

	avatarPath := fmt.Sprintf("./assets/avatars/%s%s", avatarId, filepath.Ext(file.Filename))

	dst, err := os.Create(avatarPath)
	if err != nil {
		logging.Logger.Errorf("Failed to create a new file for the avatar: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		logging.Logger.Errorf("Failed to copy the temp avatar to the final location: %s", err.Error())
		os.Remove(avatarPath)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if strings.HasPrefix(user.AvatarUrl, config.Cfg.Api.Host) {
		os.Remove(fmt.Sprintf("./assets/avatars/%s", filepath.Base(user.AvatarUrl)))
	}

	db.DBQueries.SetUserAvatar(db.DBContext, db.SetUserAvatarParams{
		ID:        user.Id,
		AvatarUrl: fmt.Sprintf("%s/assets/avatars/%s%s", config.Cfg.Api.Host, avatarId, filepath.Ext(file.Filename)),
	})

	return nil
}