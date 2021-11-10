# geofr 

> French software

![ci_workflow](https://github.com/edijon/geofr/actions/workflows/ci.yml/badge.svg)

## Sommaire

1. [Introduction](#introduction)
2. [Utilisation](#utilisation)
    1. [Aide](#aide)
    2. [Régions](#régions)
    3. [Départements](#départements)
    4. [Communes](#communes)
2. [Compilation](#compilation)
2. [Tests](#tests)

## Introduction

Client `Go` permettant de récupérer le découpage administratif Français.
- API : https://geo.api.gouv.fr/decoupage-administratif
- Go : https://golang.org/
- Pédagogique
    > Me permet de pratiquer le langage Go.

## Utilisation

### Aide

- Obtenir de l'aide
    ```bash
    ./geofr --help
    ```

### Régions

- Découpage administratif des régions
    ```bash
    ./geofr regions
    ```
- Découpage administratif de la région 76
    ```bash
    ./geofr regions 76
    ```
- Exemple
    ```bash
    # -> % ./geofr regions 76
    NOM                                     CODE                                    
    Occitanie                               76 
    ```

### Départements

- Découpage administratif des départements
    ```bash
    ./geofr departements
    ```
- Découpage administratif du département 66
    ```bash
    ./geofr departements 66
    ```

- Exemple : 
    ```bash
    # ./geofr departements 66
    NOM                      CODE                     CODEREGION               
    Pyrénées-Orientales      66                       76            
    ```

### Communes

- Découpage administratif des communes
    ```bash
    ./geofr departements
    ```
- Découpage administratif de la commune Saint-Estève
    ```bash
    ./geofr communes 66172
    ```
    ```bash
    # ./geofr communes 66172
    NOM                      CODE                     CODEDEPARTEMENT          CODEREGION               CODESPOSTAUX             POPULATION               
    Saint-Estève             66172                    66                       76                       66240                    11841     
    ```

## Compilation

- Depuis le dossier du projet 
    ```bash
    go build
    ```

## Tests

- Depuis le dossier du projet 
    ```bash
    go test ./...
    ```
