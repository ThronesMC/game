package skin

import (
	"encoding/json"
	"fmt"
	"github.com/thronesmc/game/game/utils/randomskins/config"
	"github.com/thronesmc/game/game/utils/randomskins/utils"
	"image"
	"image/draw"
	"image/png"
	"io"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	Layer0SkinFolder = "layer_0"
	Layer1SkinFolder = "layer_1"

	BaseSkinFolder     = "/base"
	HeadSkinFolder     = "/head"
	BodySkinFolder     = "/body"
	LeftArmSkinFolder  = "/left_arm"
	RightArmSkinFolder = "/righ_arm"
	LeftLegSkinFolder  = "/left_leg"
	RightLegSkinFolder = "/right_leg"
)

type Manager struct {
	Config config.File

	RandomizerFolderPath string
	SkinPath             string

	BasePart     Part
	HeadPart     Part
	BodyPart     Part
	LeftArmPart  Part
	RightArmPart Part
	LeftLegPart  Part
	RightLegPart Part
}

type Part struct {
	RootPath   string
	Layer0Path string
	Layer1Path string
}

func (sm *Manager) setupSkinParts() error {
	necessaryParts := []*Part{&sm.BasePart, &sm.HeadPart, &sm.BodyPart, &sm.LeftArmPart, &sm.RightArmPart, &sm.LeftLegPart, &sm.RightLegPart}

	for _, part := range necessaryParts {
		part.Layer0Path = fmt.Sprintf("%s/%s", part.RootPath, Layer0SkinFolder)
		part.Layer1Path = fmt.Sprintf("%s/%s", part.RootPath, Layer1SkinFolder)

		if _, err := os.Stat(part.RootPath); os.IsNotExist(err) {
			err := os.Mkdir(part.RootPath, 0755)
			if err != nil {
				return fmt.Errorf("error creating folder: %w", err)
			}
		}

		if _, err := os.Stat(part.Layer0Path); os.IsNotExist(err) {
			err := os.Mkdir(part.Layer0Path, 0755)
			if err != nil {
				return fmt.Errorf("error creating folder: %w", err)
			}
		}

		if _, err := os.Stat(part.Layer1Path); os.IsNotExist(err) {
			err := os.Mkdir(part.Layer1Path, 0755)
			if err != nil {
				return fmt.Errorf("error creating folder: %w", err)
			}
		}
	}

	return nil
}

func SetupSkinManager() (skinManager Manager, err error) {
	// the most important thing is the config file,
	// add flag to config file json
	// - stop asking for configuration via user input?

	// if config file does not exist users can generate one
	// just use the flags, and it will generate the config filepath
	//
	// flags:
	// config, skin-dir, randomize-dir => skinDir and randomizeDir you can create a config folder
	// if the config is set you can overwrite it with skin dir and randomizerDir, this should
	// save the new configuration in the config file
	//
	// recolor generation configs are experimental, so I will not add flags for it for now

	file, err := os.Open(path.Join(".", "skins", "config.json"))
	if err != nil {
		return skinManager, err
	}

	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return skinManager, err
	}

	err = json.Unmarshal(bytes, &skinManager.Config)
	if err != nil {
		return skinManager, fmt.Errorf("your config file is not correct")
	}

	if skinManager.Config.RandomizerFolder == "" {
		return skinManager, fmt.Errorf("randomizer folder path is missing")
	}

	if skinManager.Config.EditableSkin == "" {
		return skinManager, fmt.Errorf("editable skin path is missing")
	}

	skinManager.RandomizerFolderPath = skinManager.Config.RandomizerFolder

	skinManager.SkinPath = skinManager.Config.EditableSkin

	skinManager.BasePart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, BaseSkinFolder)
	skinManager.HeadPart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, HeadSkinFolder)
	skinManager.BodyPart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, BodySkinFolder)
	skinManager.LeftArmPart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, LeftArmSkinFolder)
	skinManager.RightArmPart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, RightArmSkinFolder)
	skinManager.LeftLegPart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, LeftLegSkinFolder)
	skinManager.RightLegPart.RootPath = fmt.Sprintf("%s%s", skinManager.RandomizerFolderPath, RightLegSkinFolder)

	err = skinManager.setupSkinParts()
	if err != nil {
		return skinManager, err
	}

	return skinManager, nil
}

func (sm *Manager) GenerateSkin(id int) error {
	//TODO: change to use only parts in the config

	//Order
	parts := []Part{sm.BasePart, sm.HeadPart, sm.BodyPart, sm.LeftArmPart, sm.RightArmPart, sm.LeftLegPart, sm.RightLegPart}

	finalMix := []string{}

	for _, part := range parts {
		dirLayer0 := part.Layer0Path
		dirLayer1 := part.Layer1Path

		filesLayer0, err := os.ReadDir(dirLayer0)
		if err != nil {
			return err
		}
		filesLayer1, err := os.ReadDir(dirLayer1)
		if err != nil {
			return err
		}

		var layer0Files []string
		var layer1Files []string

		for _, file := range filesLayer0 {
			if filepath.Ext(file.Name()) == ".png" {
				skinPath := dirLayer0 + "/" + file.Name()
				err = utils.VerifySkin(skinPath)
				if err != nil {
					return fmt.Errorf("layer 0 is in the wrong size: %v (%s)", err, file.Name())
				}
				layer0Files = append(layer0Files, skinPath)
			}
		}
		for _, file := range filesLayer1 {
			if filepath.Ext(file.Name()) == ".png" {
				skinPath := dirLayer1 + "/" + file.Name()
				err = utils.VerifySkin(skinPath)
				if err != nil {
					return fmt.Errorf("layer 1 is in the wrong size: %v (%s)", err, file.Name())
				}
				layer1Files = append(layer1Files, skinPath)
			}
		}

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		if len(layer0Files) != 0 {
			randomIndexLayer0 := r.Intn(len(layer0Files))
			randomLayer0Part := layer0Files[randomIndexLayer0]

			finalMix = append(finalMix, randomLayer0Part)
		}
		if len(layer1Files) != 0 {
			randomIndexLayer1 := r.Intn(len(layer1Files))
			randomLayer1Part := layer1Files[randomIndexLayer1]

			finalMix = append(finalMix, randomLayer1Part)
		}
	}

	skin, err := utils.LoadImage(finalMix[0]) // -> base skin layer 0
	if err != nil {
		return fmt.Errorf("error loading base skin image: %v", err)
	}

	for _, layerPath := range finalMix {
		currentSkinPart, err := utils.LoadImage(layerPath)
		if err != nil {
			return err
		}

		tempImage := image.NewRGBA(skin.Bounds())
		draw.Draw(tempImage, skin.Bounds(), skin, image.Point{}, draw.Over)
		draw.Draw(tempImage, currentSkinPart.Bounds().Add(image.Point{0, 0}), currentSkinPart, image.Point{}, draw.Over)

		skin = tempImage
	}

	outputFile, err := os.Create(fmt.Sprintf("%v/skin_%v.png", sm.Config.EditableSkin, id))
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, skin)
	if err != nil {
		return fmt.Errorf("error saving final image: %v", err)
	}

	return nil
}
