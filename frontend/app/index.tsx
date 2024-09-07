import { View, Text, Image } from "react-native";
import { Link, useNavigation } from "expo-router";
import { SafeAreaView } from "react-native-safe-area-context";
import { Button, ButtonText } from "@/components/ui/button";
import { AlertDialog, AlertDialogBackdrop, AlertDialogBody, AlertDialogContent, AlertDialogFooter, AlertDialogHeader } from "@/components/ui/alert-dialog";
import { Heading } from "@/components/ui/heading";
import React from "react";
import { LogBox } from 'react-native';



export default function EntryScreen() {
  LogBox.ignoreAllLogs(); // Desativa todos os alertas amarelos


  const [showAlertDialog, setShowAlertDialog] = React.useState(false)
  const handleClose = () => setShowAlertDialog(false)
  const navigation = useNavigation();

  return (
    <>
      <AlertDialog isOpen={showAlertDialog} onClose={handleClose} size="md">
        <AlertDialogBackdrop />
        <AlertDialogContent>
          <AlertDialogHeader>
            <Heading className="text-typography-950 font-semibold" size="md">
              Bem vindo ao Pernanbuco Mulher
            </Heading>
          </AlertDialogHeader>
          <AlertDialogBody className="mt-3 mb-4">
            <Text size="sm">
              Personalize o nome do app e a logo para uma maior privacidade
            </Text>
          </AlertDialogBody>
          <AlertDialogFooter className="">
            <Button
              variant="outline"
              action="secondary"
              onPress={() => {
                setShowAlertDialog(false);
                navigation.navigate("(tabs)");
              }}
              size="sm"
            >
              <ButtonText>Não, obrigado</ButtonText>
            </Button>
            <Button size="sm" onPress={handleClose}>
              <ButtonText>Ir</ButtonText>
            </Button>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <SafeAreaView
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#f5f5f5",
      }}
    >
      <View className="flex flex-1 items-center justify-center w-full p-5">
        <Image
          source={{
            uri: "https://www.pe.gov.br/wp-content/uploads/2023/08/GovPERGB.png",
          }}
          style={{ width: 200, height: 200, objectFit: "contain" }}
        />

        <Text className="text-center text-2xl bold">Pernambuco Mulher</Text>

        <View className="mt-5 mb-5">
          <Text className="text-center color-gray-400">
            Tudo para a cidadã na palma da mão
          </Text>
        </View>
        {/* Botão de Logar que leva para as tabs */}
        <Link href="/(tabs)" asChild>
          <Button
            size="md"
            variant="solid"
            action="primary"
            style={{
              backgroundColor: "#1351b4",
              width: "100%",
              marginBottom: 10,
            }}
          >
            <ButtonText>Logar com GovBR</ButtonText>
          </Button>
        </Link>
        {/* Botão de Registrar */}
        {/* <Link href="/(tabs)" asChild className="mt-2"> */}
          <Button
            size="md"
            variant="solid"
            action="primary"
            style={{
              backgroundColor: "#6366f1",
              width: "100%",
            }}
            onPress={() => setShowAlertDialog(true)}
          >
            <ButtonText>Logar de maneira anonima</ButtonText>
          </Button>
        {/* </Link> */}
      </View>
    </SafeAreaView>
    </>
    
  );
}
